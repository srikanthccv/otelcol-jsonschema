package components

import (
	"github.com/knadh/koanf/v2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/datadogprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/remoteobserverprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanmetricsprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"
	"go.opentelemetry.io/collector/processor/batchprocessor"
	"go.opentelemetry.io/collector/processor/memorylimiterprocessor"

	"github.com/srikanthccv/otel/collector/schema/jsonschema"
)

func Processors(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {

	attributes := KoanfLoad(r.Reflect(&attributesprocessor.Config{}).MarshalJSON())
	batch := KoanfLoad(r.Reflect(&batchprocessor.Config{}).MarshalJSON())
	cumulativetodelta := KoanfLoad(r.Reflect(&cumulativetodeltaprocessor.Config{}).MarshalJSON())
	datadog := KoanfLoad(r.Reflect(&datadogprocessor.Config{}).MarshalJSON())
	deltatorate := KoanfLoad(r.Reflect(&deltatorateprocessor.Config{}).MarshalJSON())
	filter := KoanfLoad(r.Reflect(&filterprocessor.Config{}).MarshalJSON())
	groupbyattrs := KoanfLoad(r.Reflect(&groupbyattrsprocessor.Config{}).MarshalJSON())
	groupbytrace := KoanfLoad(r.Reflect(&groupbytraceprocessor.Config{}).MarshalJSON())
	k8sattributes := KoanfLoad(r.Reflect(&k8sattributesprocessor.Config{}).MarshalJSON())
	memorylimiter := KoanfLoad(r.Reflect(&memorylimiterprocessor.Config{}).MarshalJSON())
	metricsgeneration := KoanfLoad(r.Reflect(&metricsgenerationprocessor.Config{}).MarshalJSON())
	metricstransform := KoanfLoad(r.Reflect(&metricstransformprocessor.Config{}).MarshalJSON())
	probabilisticsampler := KoanfLoad(r.Reflect(&probabilisticsamplerprocessor.Config{}).MarshalJSON())
	redaction := KoanfLoad(r.Reflect(&redactionprocessor.Config{}).MarshalJSON())
	remoteobserver := KoanfLoad(r.Reflect(&remoteobserverprocessor.Config{}).MarshalJSON())
	resourcedetection := KoanfLoad(r.Reflect(&resourcedetectionprocessor.Config{}).MarshalJSON())
	resource := KoanfLoad(r.Reflect(&resourceprocessor.Config{}).MarshalJSON())
	routing := KoanfLoad(r.Reflect(&routingprocessor.Config{}).MarshalJSON())
	schemaprocessor := KoanfLoad(r.Reflect(&schemaprocessor.Config{}).MarshalJSON())
	servicegraph := KoanfLoad(r.Reflect(&servicegraphprocessor.Config{}).MarshalJSON())
	spanmetrics := KoanfLoad(r.Reflect(&spanmetricsprocessor.Config{}).MarshalJSON())
	span := KoanfLoad(r.Reflect(&spanprocessor.Config{}).MarshalJSON())
	tailsampling := KoanfLoad(r.Reflect(&tailsamplingprocessor.Config{}).MarshalJSON())
	transform := KoanfLoad(r.Reflect(&transformprocessor.Config{}).MarshalJSON())

	processors := []*koanf.Koanf{
		attributes,
		batch,
		cumulativetodelta,
		datadog,
		deltatorate,
		filter,
		groupbyattrs,
		groupbytrace,
		k8sattributes,
		memorylimiter,
		metricsgeneration,
		metricstransform,
		probabilisticsampler,
		redaction,
		remoteobserver,
		resourcedetection,
		resource,
		routing,
		schemaprocessor,
		servicegraph,
		spanmetrics,
		span,
		tailsampling,
		transform,
	}

	processorsTopLevel := jsonschema.Schema{
		MinProperties: 1,
		Type:          "object",
		PatternProperties: map[string]*jsonschema.Schema{
			"^attributes(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.attributesprocessor.Config"},
			"^batch(\\/.+)?$":                {Ref: "#/$defs/go.opentelemetry.io.collector.processor.batchprocessor.Config"},
			"^cumulativetodelta(\\/.+)?$":    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.cumulativetodeltaprocessor.Config"},
			"^datadog(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.datadogprocessor.Config"},
			"^deltatorate(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.deltatorateprocessor.Config"},
			"^filter(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.filterprocessor.Config"},
			"^groupbyattrs(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.groupbyattrsprocessor.Config"},
			"^groupbytrace(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.groupbytraceprocessor.Config"},
			"^k8sattributes(\\/.+)?$":        {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.k8sattributesprocessor.Config"},
			"^memorylimiter(\\/.+)?$":        {Ref: "#/$defs/go.opentelemetry.io.collector.processor.memorylimiterprocessor.Config"},
			"^metricsgeneration(\\/.+)?$":    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.metricsgenerationprocessor.Config"},
			"^metricstransform(\\/.+)?$":     {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.metricstransformprocessor.Config"},
			"^probabilisticsampler(\\/.+)?$": {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.probabilisticsamplerprocessor.Config"},
			"^redaction(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.redactionprocessor.Config"},
			"^remoteobserver(\\/.+)?$":       {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.remoteobserverprocessor.Config"},
			"^resourcedetection(\\/.+)?$":    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.resourcedetectionprocessor.Config"},
			"^resource(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.resourceprocessor.Config"},
			"^routing(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.routingprocessor.Config"},
			"^schemaprocessor(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.schemaprocessor.Config"},
			"^servicegraph(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.servicegraphprocessor.Config"},
			"^spanmetrics(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.spanmetricsprocessor.Config"},
			"^span(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.spanprocessor.Config"},
			"^tailsampling(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.tailsamplingprocessor.Config"},
			"^transform(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.transformprocessor.Config"},
		},
	}
	return processors, &processorsTopLevel
}
