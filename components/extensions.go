package components

import (
	"github.com/knadh/koanf/v2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/awsproxy"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarder"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/oidcauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"
	"go.opentelemetry.io/collector/extension/ballastextension"
	"go.opentelemetry.io/collector/extension/zpagesextension"

	"github.com/srikanthccv/otel/collector/schema/jsonschema"
)

func Extensions(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {
	asapauth := KoanfLoad(r.Reflect(&asapauthextension.Config{}).MarshalJSON())
	awsproxy := KoanfLoad(r.Reflect(&awsproxy.Config{}).MarshalJSON())
	ballast := KoanfLoad(r.Reflect(&ballastextension.Config{}).MarshalJSON())
	basicauthextension := KoanfLoad(r.Reflect(&basicauthextension.Config{}).MarshalJSON())
	bearertokenauthextension := KoanfLoad(r.Reflect(&bearertokenauthextension.Config{}).MarshalJSON())
	dbstorage := KoanfLoad(r.Reflect(&dbstorage.Config{}).MarshalJSON())
	filestorage := KoanfLoad(r.Reflect(&filestorage.Config{}).MarshalJSON())
	headerssetter := KoanfLoad(r.Reflect(&headerssetterextension.Config{}).MarshalJSON())
	healthcheckextension := KoanfLoad(r.Reflect(&healthcheckextension.Config{}).MarshalJSON())
	httpforwarder := KoanfLoad(r.Reflect(&httpforwarder.Config{}).MarshalJSON())
	jaegerremotesampling := KoanfLoad(r.Reflect(&jaegerremotesampling.Config{}).MarshalJSON())
	oauth2clientauthextension := KoanfLoad(r.Reflect(&oauth2clientauthextension.Config{}).MarshalJSON())
	dockerobserver := KoanfLoad(r.Reflect(&dockerobserver.Config{}).MarshalJSON())
	ecsobserver := KoanfLoad(r.Reflect(&ecsobserver.Config{}).MarshalJSON())
	ecstaskobserver := KoanfLoad(r.Reflect(&ecstaskobserver.Config{}).MarshalJSON())
	hostobserver := KoanfLoad(r.Reflect(&hostobserver.Config{}).MarshalJSON())
	k8sobserver := KoanfLoad(r.Reflect(&k8sobserver.Config{}).MarshalJSON())
	oidcauthextension := KoanfLoad(r.Reflect(&oidcauthextension.Config{}).MarshalJSON())
	pprofextension := KoanfLoad(r.Reflect(&pprofextension.Config{}).MarshalJSON())
	sigv4auth := KoanfLoad(r.Reflect(&sigv4authextension.Config{}).MarshalJSON())
	zpages := KoanfLoad(r.Reflect(&zpagesextension.Config{}).MarshalJSON())

	extensions := []*koanf.Koanf{
		asapauth,
		awsproxy,
		ballast,
		basicauthextension,
		bearertokenauthextension,
		dbstorage,
		filestorage,
		headerssetter,
		healthcheckextension,
		httpforwarder,
		jaegerremotesampling,
		oauth2clientauthextension,
		dockerobserver,
		ecsobserver,
		ecstaskobserver,
		hostobserver,
		k8sobserver,
		oidcauthextension,
		pprofextension,
		sigv4auth,
		zpages,
	}

	extensionsTopLevel := jsonschema.Schema{
		MinProperties: 1,
		Type:          "object",
		PatternProperties: map[string]*jsonschema.Schema{
			"^asapauth(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.asapauthextension.Config"},
			"^awsproxy(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.awsproxy.Config"},
			"^ballast(\\/.+)?$":              {Ref: "#/$defs/go.opentelemetry.io.collector.extension.ballastextension.Config"},
			"^basicauth(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.basicauthextension.Config"},
			"^bearertokenauth(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.bearertokenauthextension.Config"},
			"^dbstorage(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.storage.dbstorage.Config"},
			"^filestorage(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.storage.filestorage.Config"},
			"^headerssetter(\\/.+)?$":        {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.headerssetterextension.Config"},
			"^healthcheck(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.healthcheckextension.Config"},
			"^httpforwarder(\\/.+)?$":        {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.httpforwarder.Config"},
			"^jaegerremotesampling(\\/.+)?$": {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.jaegerremotesampling.Config"},
			"^oauth2clientauth(\\/.+)?$":     {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.oauth2clientauthextension.Config"},
			"^dockerobserver(\\/.+)?$":       {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.observer.dockerobserver.Config"},
			"^ecsobserver(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.observer.ecsobserver.Config"},
			"^ecstaskobserver(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.observer.ecstaskobserver.Config"},
			"^hostobserver(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.observer.hostobserver.Config"},
			"^k8sobserver(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.observer.k8sobserver.Config"},
			"^oidcauth(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.oidcauthextension.Config"},
			"^pprof(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.pprofextension.Config"},
			"^sigv4auth(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.sigv4authextension.Config"},
			"^zpages(\\/.+)?$":               {Ref: "#/$defs/go.opentelemetry.io.collector.extension.zpagesextension.Config"},
		},
	}
	return extensions, &extensionsTopLevel
}
