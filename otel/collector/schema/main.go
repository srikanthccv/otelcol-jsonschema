package main

import (
	"fmt"
	"go/build"
	"log"
	"os"

	"github.com/iancoleman/orderedmap"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/jaegerexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/jaegerthrifthttpexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/parquetexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"
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
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanmetricsprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dotnetdiagnosticsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusexecreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zookeeperreceiver"
	"go.opentelemetry.io/collector/service"

	"github.com/srikanthccv/otel/collector/schema/jsonschema"
)

func modCacheLoc() string {
	return build.Default.GOPATH + "/pkg/mod"
}

func koanfLoad(b []byte, err error) *koanf.Koanf {
	if err != nil {
		log.Fatal(err)
	}
	var k = koanf.New(".")
	k.Load(rawbytes.Provider(b), json.Parser())
	return k
}

func debug(r *jsonschema.Reflector) {
	jaegerthrift := koanfLoad(r.Reflect(&service.Config{}).MarshalJSON())

	final := koanf.New(".")
	final.Merge(jaegerthrift)
	s, _ := final.Marshal(json.Parser())
	fmt.Println(string(s))
}

func collectorService(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {

	service := koanfLoad(r.Reflect(&service.Config{}).MarshalJSON())

	serviceSchema := jsonschema.Schema{
		Ref: "#/$defs/go.opentelemetry.io.collector.service.Config",
	}

	return []*koanf.Koanf{service}, &serviceSchema
}

func exporters(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {

	carbon := koanfLoad(r.Reflect(&carbonexporter.Config{}).MarshalJSON())
	file := koanfLoad(r.Reflect(&fileexporter.Config{}).MarshalJSON())
	jaeger := koanfLoad(r.Reflect(&jaegerexporter.Config{}).MarshalJSON())
	jaegerthrift := koanfLoad(r.Reflect(&jaegerthrifthttpexporter.Config{}).MarshalJSON())
	kafka := koanfLoad(r.Reflect(&kafkaexporter.Config{}).MarshalJSON())
	loadbalancing := koanfLoad(r.Reflect(&loadbalancingexporter.Config{}).MarshalJSON())
	opencensus := koanfLoad(r.Reflect(&opencensusexporter.Config{}).MarshalJSON())
	parquet := koanfLoad(r.Reflect(&parquetexporter.Config{}).MarshalJSON())
	prometheus := koanfLoad(r.Reflect(&prometheusexporter.Config{}).MarshalJSON())
	prometheusremotewrite := koanfLoad(r.Reflect(&prometheusremotewriteexporter.Config{}).MarshalJSON())
	pulsar := koanfLoad(r.Reflect(&pulsarexporter.Config{}).MarshalJSON())
	zipkin := koanfLoad(r.Reflect(&zipkinexporter.Config{}).MarshalJSON())

	exporters := []*koanf.Koanf{
		carbon,
		file,
		jaeger,
		jaegerthrift,
		kafka,
		loadbalancing,
		opencensus,
		parquet,
		prometheus,
		prometheusremotewrite,
		pulsar,
		zipkin,
	}

	exportersTopLevel := jsonschema.Schema{
		MinProperties: 1,
		Type:          "object",
		PatternProperties: map[string]*jsonschema.Schema{
			"^carbon(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.carbonexporter.Config"},
			"^file(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.fileexporter.Config"},
			"^jaeger(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.jaegerexporter.Config"},
			"^jaegerthrift(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.jaegerthrifthttpexporter.Config"},
			"^kafka(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.kafkaexporter.Config"},
			"^loadbalancing(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.loadbalancingexporter.Config"},
			"^opencensus(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.opencensusexporter.Config"},
			"^parquet(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.parquetexporter.Config"},
			"^prometheus(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.prometheusexporter.Config"},
			"^prometheusremotewrite(\\/.+)?$": {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.prometheusremotewriteexporter.Config"},
			"^pulsar(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.pulsarexporter.Config"},
			"^zipkin(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.zipkinexporter.Config"},
		},
	}
	return exporters, &exportersTopLevel
}

func extensions(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {

	basicauth := koanfLoad(r.Reflect(&basicauthextension.Config{}).MarshalJSON())
	bearertokenauth := koanfLoad(r.Reflect(&bearertokenauthextension.Config{}).MarshalJSON())
	healthcheck := koanfLoad(r.Reflect(&healthcheckextension.Config{}).MarshalJSON())
	httpforwarder := koanfLoad(r.Reflect(&httpforwarder.Config{}).MarshalJSON())
	jaegerremotesampling := koanfLoad(r.Reflect(&jaegerremotesampling.Config{}).MarshalJSON())
	oauth2clientauth := koanfLoad(r.Reflect(&oauth2clientauthextension.Config{}).MarshalJSON())
	dockerobserver := koanfLoad(r.Reflect(&dockerobserver.Config{}).MarshalJSON())
	ecsobserver := koanfLoad(r.Reflect(&ecsobserver.Config{}).MarshalJSON())
	ecstaskobserver := koanfLoad(r.Reflect(&ecstaskobserver.Config{}).MarshalJSON())
	hostobserver := koanfLoad(r.Reflect(&hostobserver.Config{}).MarshalJSON())
	k8sobserver := koanfLoad(r.Reflect(&k8sobserver.Config{}).MarshalJSON())
	oidcauth := koanfLoad(r.Reflect(&oidcauthextension.Config{}).MarshalJSON())
	pprof := koanfLoad(r.Reflect(&pprofextension.Config{}).MarshalJSON())

	extensions := []*koanf.Koanf{
		basicauth,
		bearertokenauth,
		healthcheck,
		httpforwarder,
		jaegerremotesampling,
		oauth2clientauth,
		dockerobserver,
		ecsobserver,
		ecstaskobserver,
		hostobserver,
		k8sobserver,
		oidcauth,
		pprof,
	}

	extensionsTopLevel := jsonschema.Schema{
		MinProperties: 1,
		Type:          "object",
		PatternProperties: map[string]*jsonschema.Schema{
			"^basicauth(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.basicauthextension.Config"},
			"^bearertokenauth(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.extension.bearertokenauthextension.Config"},
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
		},
	}
	return extensions, &extensionsTopLevel
}

func receivers(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {

	apache := koanfLoad(r.Reflect(&apachereceiver.Config{}).MarshalJSON())
	awscloudwatch := koanfLoad(r.Reflect(&awscloudwatchreceiver.Config{}).MarshalJSON())
	awscontainerinsight := koanfLoad(r.Reflect(&awscontainerinsightreceiver.Config{}).MarshalJSON())
	awsecscontainermetrics := koanfLoad(r.Reflect(&awsecscontainermetricsreceiver.Config{}).MarshalJSON())
	awsxray := koanfLoad(r.Reflect(&awsxrayreceiver.Config{}).MarshalJSON())
	carbon := koanfLoad(r.Reflect(&carbonreceiver.Config{}).MarshalJSON())
	collectd := koanfLoad(r.Reflect(&collectdreceiver.Config{}).MarshalJSON())
	couchdb := koanfLoad(r.Reflect(&couchdbreceiver.Config{}).MarshalJSON())
	dockerstats := koanfLoad(r.Reflect(&dockerstatsreceiver.Config{}).MarshalJSON())
	dotnetdiagnostics := koanfLoad(r.Reflect(&dotnetdiagnosticsreceiver.Config{}).MarshalJSON())
	elasticsearch := koanfLoad(r.Reflect(&elasticsearchreceiver.Config{}).MarshalJSON())
	// filelog := koanfLoad(r.Reflect(&filelogreceiver.FileLogConfig{}).MarshalJSON())
	fluentforward := koanfLoad(r.Reflect(&fluentforwardreceiver.Config{}).MarshalJSON())
	// TODO: https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/6816bc1448c01f168bf7cf07bb0cea3256b19a66/receiver/hostmetricsreceiver/config.go#L25
	// hostmetrics := koanfLoad(r.Reflect(&hostmetricsreceiver.Config{}).MarshalJSON())
	httpcheck := koanfLoad(r.Reflect(&httpcheckreceiver.Config{}).MarshalJSON())
	jaeger := koanfLoad(r.Reflect(&jaegerreceiver.Config{}).MarshalJSON())
	jmx := koanfLoad(r.Reflect(&jmxreceiver.Config{}).MarshalJSON())
	// journald := koanfLoad(r.Reflect(&journaldreceiver.JournaldConfig{}).MarshalJSON())
	k8scluster := koanfLoad(r.Reflect(&k8sclusterreceiver.Config{}).MarshalJSON())
	k8sevents := koanfLoad(r.Reflect(&k8seventsreceiver.Config{}).MarshalJSON())
	k8sobjects := koanfLoad(r.Reflect(&k8sobjectsreceiver.Config{}).MarshalJSON())
	kafkametrics := koanfLoad(r.Reflect(&kafkametricsreceiver.Config{}).MarshalJSON())
	kafka := koanfLoad(r.Reflect(&kafkareceiver.Config{}).MarshalJSON())
	kubeletstats := koanfLoad(r.Reflect(&kubeletstatsreceiver.Config{}).MarshalJSON())
	memcached := koanfLoad(r.Reflect(&memcachedreceiver.Config{}).MarshalJSON())
	mongodbatlas := koanfLoad(r.Reflect(&mongodbatlasreceiver.Config{}).MarshalJSON())
	mongodb := koanfLoad(r.Reflect(&mongodbreceiver.Config{}).MarshalJSON())
	mysql := koanfLoad(r.Reflect(&mysqlreceiver.Config{}).MarshalJSON())
	nginx := koanfLoad(r.Reflect(&nginxreceiver.Config{}).MarshalJSON())
	opencensus := koanfLoad(r.Reflect(&opencensusreceiver.Config{}).MarshalJSON())
	oracledb := koanfLoad(r.Reflect(&oracledbreceiver.Config{}).MarshalJSON())
	// otlpjsonfile := koanfLoad(r.Reflect(&otlpjsonfilereceiver.Config{}).MarshalJSON())
	podman := koanfLoad(r.Reflect(&podmanreceiver.Config{}).MarshalJSON())
	postgresql := koanfLoad(r.Reflect(&postgresqlreceiver.Config{}).MarshalJSON())
	prometheusexec := koanfLoad(r.Reflect(&prometheusexecreceiver.Config{}).MarshalJSON())
	prometheus := koanfLoad(r.Reflect(&prometheusreceiver.Config{}).MarshalJSON())
	pulsar := koanfLoad(r.Reflect(&pulsarreceiver.Config{}).MarshalJSON())
	rabbitmq := koanfLoad(r.Reflect(&rabbitmqreceiver.Config{}).MarshalJSON())
	receivercreator := koanfLoad(r.Reflect(&receivercreator.Config{}).MarshalJSON())
	redis := koanfLoad(r.Reflect(&redisreceiver.Config{}).MarshalJSON())
	simpleprometheus := koanfLoad(r.Reflect(&simpleprometheusreceiver.Config{}).MarshalJSON())
	sqlquery := koanfLoad(r.Reflect(&sqlqueryreceiver.Config{}).MarshalJSON())
	sqlserver := koanfLoad(r.Reflect(&sqlserverreceiver.Config{}).MarshalJSON())
	statsd := koanfLoad(r.Reflect(&statsdreceiver.Config{}).MarshalJSON())
	// syslog := koanfLoad(r.Reflect(&syslogreceiver.SysLogConfig{}).MarshalJSON())
	// tcplog := koanfLoad(r.Reflect(&tcplogreceiver.TCPLogConfig{}).MarshalJSON())
	// udplog := koanfLoad(r.Reflect(&udplogreceiver.UDPLogConfig{}).MarshalJSON())
	zipkin := koanfLoad(r.Reflect(&zipkinreceiver.Config{}).MarshalJSON())
	zookeeper := koanfLoad(r.Reflect(&zookeeperreceiver.Config{}).MarshalJSON())

	receivers := []*koanf.Koanf{
		apache,
		awscloudwatch,
		awscontainerinsight,
		awsecscontainermetrics,
		awsxray,
		carbon,
		collectd,
		couchdb,
		dockerstats,
		dotnetdiagnostics,
		elasticsearch,
		// filelog,
		fluentforward,
		// hostmetrics,
		httpcheck,
		jaeger,
		jmx,
		// journald,
		k8scluster,
		k8sevents,
		k8sobjects,
		kafkametrics,
		kafka,
		kubeletstats,
		memcached,
		mongodbatlas,
		mongodb,
		mysql,
		nginx,
		opencensus,
		oracledb,
		// otlpjsonfile,
		podman,
		postgresql,
		prometheusexec,
		prometheus,
		pulsar,
		rabbitmq,
		receivercreator,
		redis,
		simpleprometheus,
		sqlquery,
		sqlserver,
		statsd,
		// syslog,
		// tcplog,
		// udplog,
		zipkin,
		zookeeper,
	}

	receiversTopLevel := jsonschema.Schema{
		MinProperties: 1,
		Type:          "object",
		PatternProperties: map[string]*jsonschema.Schema{
			"^apache(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.apachereceiver.Config"},
			"^awscloudwatch(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awscloudwatchreceiver.Config"},
			"^awscontainerinsight(\\/.+)?$":    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awscontainerinsightreceiver.Config"},
			"^awsecscontainermetrics(\\/.+)?$": {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awsecscontainermetricsreceiver.Config"},
			"^awsxray(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awsxrayreceiver.Config"},
			"^carbon(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.carbonreceiver.Config"},
			"^collectd(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.collectdreceiver.Config"},
			"^couchdb(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.couchdbreceiver.Config"},
			"^dockerstats(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.dockerstatsreceiver.Config"},
			"^dotnetdiagnostics(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.dotnetdiagnosticsreceiver.Config"},
			"^elasticsearch(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.elasticsearchreceiver.Config"},
			"^filelog(\\/.+)?$":                { /*Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.filelogreceiver.FileLogConfig"*/ },
			"^fluentforward(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.fluentforwardreceiver.Config"},
			"^hostmetrics(\\/.+)?$":            { /*Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.hostmetricsreceiver.Config"*/ },
			"^httpcheck(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.httpcheckreceiver.Config"},
			"^jaeger(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.jaegerreceiver.Config"},
			"^jmx(\\/.+)?$":                    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.jmxreceiver.Config"},
			"^journald(\\/.+)?$":               { /*Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.journaldreceiver.JournaldConfig"*/ },
			"^k8scluster(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.k8sclusterreceiver.Config"},
			"^k8sevents(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.k8seventsreceiver.Config"},
			"^k8sobjects(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.k8sobjectsreceiver.Config"},
			"^kafkametrics(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.kafkametricsreceiver.Config"},
			"^kafka(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.kafkareceiver.Config"},
			"^kubeletstats(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.kubeletstatsreceiver.Config"},
			"^memcached(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.memcachedreceiver.Config"},
			"^mongodbatlas(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.mongodbatlasreceiver.Config"},
			"^mongodb(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.mongodbreceiver.Config"},
			"^mysql(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.mysqlreceiver.Config"},
			"^nginx(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.nginxreceiver.Config"},
			"^opencensus(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.opencensusreceiver.Config"},
			"^oracledb(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.oracledbreceiver.Config"},
			"^otlpjsonfile(\\/.+)?$":           { /*Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.otlpjsonfilereceiver.Config"*/ },
			"^podman(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.podmanreceiver.Config"},
			"^postgresql(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.postgresqlreceiver.Config"},
			"^prometheusexec(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.prometheusexecreceiver.Config"},
			"^prometheus(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.prometheusreceiver.Config"},
			"^pulsar(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.pulsarreceiver.Config"},
			"^rabbitmq(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.rabbitmqreceiver.Config"},
			"^receivercreator(\\/.+)?$":        {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.receivercreator.Config"},
			"^redis(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.redisreceiver.Config"},
			"^simpleprometheus(\\/.+)?$":       {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.simpleprometheusreceiver.Config"},
			"^sqlquery(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.sqlqueryreceiver.Config"},
			"^sqlserver(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.sqlserverreceiver.Config"},
			"^statsd(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.statsdreceiver.Config"},
			"^syslog(\\/.+)?$":                 { /*Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.syslogreceiver.SysLogConfig"*/ },
			"^tcplog(\\/.+)?$":                 { /*Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.tcplogreceiver.TCPLogConfig"*/ },
			"^udplog(\\/.+)?$":                 { /*Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.udplogreceiver.UDPLogConfig"*/ },
			"^zipkin(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.zipkinreceiver.Config"},
			"^zookeeper(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.zookeeperreceiver.Config"},
		},
	}
	return receivers, &receiversTopLevel
}

func processors(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {
	attribute := koanfLoad(r.Reflect(&attributesprocessor.Config{}).MarshalJSON())
	cumulativetodelta := koanfLoad(r.Reflect(&cumulativetodeltaprocessor.Config{}).MarshalJSON())
	deltatorate := koanfLoad(r.Reflect(&deltatorateprocessor.Config{}).MarshalJSON())
	filter := koanfLoad(r.Reflect(&filterprocessor.Config{}).MarshalJSON())
	groupbyattrs := koanfLoad(r.Reflect(&groupbyattrsprocessor.Config{}).MarshalJSON())
	groupbytrace := koanfLoad(r.Reflect(&groupbytraceprocessor.Config{}).MarshalJSON())
	k8sattributes := koanfLoad(r.Reflect(&k8sattributesprocessor.Config{}).MarshalJSON())
	// logstransform := koanfLoad(r.Reflect(&logstransformprocessor.Config{}).MarshalJSON())
	metricsgeneration := koanfLoad(r.Reflect(&metricsgenerationprocessor.Config{}).MarshalJSON())
	metricstransform := koanfLoad(r.Reflect(&metricstransformprocessor.Config{}).MarshalJSON())
	probabilisticsampler := koanfLoad(r.Reflect(&probabilisticsamplerprocessor.Config{}).MarshalJSON())
	redaction := koanfLoad(r.Reflect(&redactionprocessor.Config{}).MarshalJSON())
	resourcedetection := koanfLoad(r.Reflect(&resourcedetectionprocessor.Config{}).MarshalJSON())
	resource := koanfLoad(r.Reflect(&resourceprocessor.Config{}).MarshalJSON())
	routing := koanfLoad(r.Reflect(&routingprocessor.Config{}).MarshalJSON())
	schema := koanfLoad(r.Reflect(&schemaprocessor.Config{}).MarshalJSON())
	servicegraph := koanfLoad(r.Reflect(&servicegraphprocessor.Config{}).MarshalJSON())
	spanmetrics := koanfLoad(r.Reflect(&spanmetricsprocessor.Config{}).MarshalJSON())
	span := koanfLoad(r.Reflect(&spanprocessor.Config{}).MarshalJSON())
	tailsampling := koanfLoad(r.Reflect(&tailsamplingprocessor.Config{}).MarshalJSON())
	transform := koanfLoad(r.Reflect(&transformprocessor.Config{}).MarshalJSON())

	processors := []*koanf.Koanf{
		attribute,
		cumulativetodelta,
		deltatorate,
		filter,
		groupbyattrs,
		groupbytrace,
		k8sattributes,
		// logstransform,
		metricsgeneration,
		metricstransform,
		probabilisticsampler,
		redaction,
		resourcedetection,
		resource,
		routing,
		schema,
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
			"^cumulativetodelta(\\/.+)?$":    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.cumulativetodeltaprocessor.Config"},
			"^deltatorate(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.deltatorateprocessor.Config"},
			"^filter(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.filterprocessor.Config"},
			"^groupbyattrs(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.groupbyattrsprocessor.Config"},
			"^groupbytrace(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.groupbytraceprocessor.Config"},
			"^k8sattributes(\\/.+)?$":        {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.k8sattributesprocessor.Config"},
			"^logstransform(\\/.+)?$":        { /*Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.logstransformprocessor.Config"*/ },
			"^metricsgeneration(\\/.+)?$":    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.metricsgenerationprocessor.Config"},
			"^metricstransform(\\/.+)?$":     {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.metricstransformprocessor.Config"},
			"^probabilisticsampler(\\/.+)?$": {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.probabilisticsamplerprocessor.Config"},
			"^redaction(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.redactionprocessor.Config"},
			"^resourcedetection(\\/.+)?$":    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.resourcedetectionprocessor.Config"},
			"^resource(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.resourceprocessor.Config"},
			"^routing(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.routingprocessor.Config"},
			"^schema(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.schemaprocessor.Config"},
			"^servicegraph(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.servicegraphprocessor.Config"},
			"^spanmetrics(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.spanmetricsprocessor.Config"},
			"^span(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.spanprocessor.Config"},
			"^tailsampling(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.tailsamplingprocessor.Config"},
			"^transform(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.processor.transformprocessor.Config"},
		},
	}
	return processors, &processorsTopLevel
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

	// debug(r)
	// return

	exporters, exportersTopLevel := exporters(r)
	extensions, extensionsTopLevel := extensions(r)
	receivers, receiversTopLevel := receivers(r)
	processors, processorsTopLevel := processors(r)

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

	final.Merge(koanfLoad(mainSchema.MarshalJSON()))
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
