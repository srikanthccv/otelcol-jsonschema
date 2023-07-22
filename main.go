package main

import (
	"fmt"
	"go/build"
	"log"
	"os"

	"github.com/iancoleman/orderedmap"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/v2"

	"go.opentelemetry.io/collector/service"

	"github.com/srikanthccv/otel/collector/schema/components"
	"github.com/srikanthccv/otel/collector/schema/jsonschema"
)

func modCacheLoc() string {
	return build.Default.GOPATH + "/pkg/mod"
}

func debug(r *jsonschema.Reflector) {
	sCfg := components.KoanfLoad(r.Reflect(&service.Config{}).MarshalJSON())

	final := koanf.New(".")
	final.Merge(sCfg)
	s, _ := final.Marshal(json.Parser())
	fmt.Println(string(s))
}

func collectorService(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {

	service := components.KoanfLoad(r.Reflect(&service.Config{}).MarshalJSON())

	serviceSchema := jsonschema.Schema{
		Ref: "#/$defs/go.opentelemetry.io.collector.service.Config",
	}

	return []*koanf.Koanf{service}, &serviceSchema
}

func main() {

	r := new(jsonschema.Reflector)
	r.Anonymous = true
	if err := r.AddGoComments("", fmt.Sprintf("%s/go.opentelemetry.io/collector", modCacheLoc())); err != nil {
		log.Print("error adding go comments for core repo")
	}
	if err := r.AddGoComments("", fmt.Sprintf("%s/github.com/open-telemetry/opentelemetry-collector-contrib", modCacheLoc())); err != nil {
		log.Print("error adding go comments for contrib repo")
	}
	final := koanf.New(".")

	exporters, exportersTopLevel := components.Exporters(r)
	extensions, extensionsTopLevel := components.Extensions(r)
	receivers, receiversTopLevel := components.Receivers(r)
	processors, processorsTopLevel := components.Processors(r)

	service, serviceSchema := collectorService(r)

	for _, e := range exporters {
		final.Merge(e)
	}
	for _, e := range extensions {
		final.Merge(e)
	}
	for _, r := range receivers {
		final.Merge(r)
	}
	for _, p := range processors {
		final.Merge(p)
	}
	for _, s := range service {
		final.Merge(s)
	}

	mainSchemaProperties := orderedmap.New()
	mainSchemaProperties.Set("receivers", receiversTopLevel)
	mainSchemaProperties.Set("exporters", exportersTopLevel)
	mainSchemaProperties.Set("processors", processorsTopLevel)
	mainSchemaProperties.Set("extensions", extensionsTopLevel)
	mainSchemaProperties.Set("service", serviceSchema)

	mainSchema := jsonschema.Schema{
		Version:    "https://json-schema.org/draft/2020-12/schema",
		Title:      "OpenTelemetry Collector Config Schema",
		Type:       "object",
		Required:   []string{"receivers", "exporters", "service"},
		Properties: mainSchemaProperties,
	}

	final.Merge(components.KoanfLoad(mainSchema.MarshalJSON()))
	final.Delete("$ref")

	s, _ := final.Marshal(json.Parser())

	// write to schema.json
	fd, err := os.Create("schema.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	fd.Write(s)
}
