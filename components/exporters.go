package components

import (
	"github.com/knadh/koanf/v2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/cassandraexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"
	dynatraceexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter/config"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/f5cloudexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlemanagedprometheusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/influxdbexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/instanaexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/jaegerexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/jaegerthrifthttpexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lokiexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/mezmoexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/parquetexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/skywalkingexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tanzuobservabilityexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter"
	"github.com/srikanthccv/otel/collector/schema/jsonschema"
)

func Exporters(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {

	alibabacloudlogservice := KoanfLoad(r.Reflect(&alibabacloudlogserviceexporter.Config{}).MarshalJSON())
	awscloudwatchlogs := KoanfLoad(r.Reflect(&awscloudwatchlogsexporter.Config{}).MarshalJSON())
	awsemf := KoanfLoad(r.Reflect(&awsemfexporter.Config{}).MarshalJSON())
	awskinesis := KoanfLoad(r.Reflect(&awskinesisexporter.Config{}).MarshalJSON())
	awsxray := KoanfLoad(r.Reflect(&awsxrayexporter.Config{}).MarshalJSON())
	awss3 := KoanfLoad(r.Reflect(&awss3exporter.Config{}).MarshalJSON())
	azuredataexplorer := KoanfLoad(r.Reflect(&azuredataexplorerexporter.Config{}).MarshalJSON())
	azuremonitor := KoanfLoad(r.Reflect(&azuremonitorexporter.Config{}).MarshalJSON())
	carbon := KoanfLoad(r.Reflect(&carbonexporter.Config{}).MarshalJSON())
	cassandra := KoanfLoad(r.Reflect(&cassandraexporter.Config{}).MarshalJSON())
	clickhouse := KoanfLoad(r.Reflect(&clickhouseexporter.Config{}).MarshalJSON())
	coralogix := KoanfLoad(r.Reflect(&coralogixexporter.Config{}).MarshalJSON())
	datadog := KoanfLoad(r.Reflect(&datadogexporter.Config{}).MarshalJSON())
	dataset := KoanfLoad(r.Reflect(&datasetexporter.Config{}).MarshalJSON())
	dynatrace := KoanfLoad(r.Reflect(&dynatraceexporter.Config{}).MarshalJSON())
	elasticsearch := KoanfLoad(r.Reflect(&elasticsearchexporter.Config{}).MarshalJSON())
	f5cloud := KoanfLoad(r.Reflect(&f5cloudexporter.Config{}).MarshalJSON())
	file := KoanfLoad(r.Reflect(&fileexporter.Config{}).MarshalJSON())
	googlecloud := KoanfLoad(r.Reflect(&googlecloudexporter.Config{}).MarshalJSON())
	googlecloudpubsub := KoanfLoad(r.Reflect(&googlecloudpubsubexporter.Config{}).MarshalJSON())
	googlemanagedprometheus := KoanfLoad(r.Reflect(&googlemanagedprometheusexporter.Config{}).MarshalJSON())
	influxdb := KoanfLoad(r.Reflect(&influxdbexporter.Config{}).MarshalJSON())
	instana := KoanfLoad(r.Reflect(&instanaexporter.Config{}).MarshalJSON())
	jaeger := KoanfLoad(r.Reflect(&jaegerexporter.Config{}).MarshalJSON())
	jaegerthrift := KoanfLoad(r.Reflect(&jaegerthrifthttpexporter.Config{}).MarshalJSON())
	kafka := KoanfLoad(r.Reflect(&kafkaexporter.Config{}).MarshalJSON())
	loadbalancing := KoanfLoad(r.Reflect(&loadbalancingexporter.Config{}).MarshalJSON())
	logicmonitor := KoanfLoad(r.Reflect(&logicmonitorexporter.Config{}).MarshalJSON())
	logzio := KoanfLoad(r.Reflect(&logzioexporter.Config{}).MarshalJSON())
	loki := KoanfLoad(r.Reflect(&lokiexporter.Config{}).MarshalJSON())
	mezmo := KoanfLoad(r.Reflect(&mezmoexporter.Config{}).MarshalJSON())
	opencensus := KoanfLoad(r.Reflect(&opencensusexporter.Config{}).MarshalJSON())
	parquet := KoanfLoad(r.Reflect(&parquetexporter.Config{}).MarshalJSON())
	prometheus := KoanfLoad(r.Reflect(&prometheusexporter.Config{}).MarshalJSON())
	prometheusremotewrite := KoanfLoad(r.Reflect(&prometheusremotewriteexporter.Config{}).MarshalJSON())
	pulsar := KoanfLoad(r.Reflect(&pulsarexporter.Config{}).MarshalJSON())
	sapm := KoanfLoad(r.Reflect(&sapmexporter.Config{}).MarshalJSON())
	sentry := KoanfLoad(r.Reflect(&sentryexporter.Config{}).MarshalJSON())
	signalfx := KoanfLoad(r.Reflect(&signalfxexporter.Config{}).MarshalJSON())
	skywalking := KoanfLoad(r.Reflect(&skywalkingexporter.Config{}).MarshalJSON())
	splunkhec := KoanfLoad(r.Reflect(&splunkhecexporter.Config{}).MarshalJSON())
	sumologic := KoanfLoad(r.Reflect(&sumologicexporter.Config{}).MarshalJSON())
	tanzuobservability := KoanfLoad(r.Reflect(&tanzuobservabilityexporter.Config{}).MarshalJSON())
	tencentcloudlogservice := KoanfLoad(r.Reflect(&tencentcloudlogserviceexporter.Config{}).MarshalJSON())
	zipkin := KoanfLoad(r.Reflect(&zipkinexporter.Config{}).MarshalJSON())

	exporters := []*koanf.Koanf{
		alibabacloudlogservice,
		awscloudwatchlogs,
		awsemf,
		awskinesis,
		awsxray,
		awss3,
		azuredataexplorer,
		azuremonitor,
		carbon,
		cassandra,
		clickhouse,
		coralogix,
		datadog,
		dataset,
		dynatrace,
		elasticsearch,
		f5cloud,
		file,
		googlecloud,
		googlecloudpubsub,
		googlemanagedprometheus,
		influxdb,
		instana,
		jaeger,
		jaegerthrift,
		kafka,
		loadbalancing,
		logicmonitor,
		logzio,
		loki,
		mezmo,
		opencensus,
		parquet,
		prometheus,
		prometheusremotewrite,
		pulsar,
		sapm,
		sentry,
		signalfx,
		skywalking,
		splunkhec,
		sumologic,
		tanzuobservability,
		tencentcloudlogservice,
		zipkin,
	}

	exportersTopLevel := jsonschema.Schema{
		MinProperties: 1,
		Type:          "object",
		PatternProperties: map[string]*jsonschema.Schema{
			"^alibabacloudlogservice(\\/.+)?$":  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.alibabacloudlogserviceexporter.Config"},
			"^awscloudwatchlogs(\\/.+)?$":       {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.awscloudwatchlogsexporter.Config"},
			"^awsemf(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.awsemfexporter.Config"},
			"^awskinesis(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.awskinesisexporter.Config"},
			"^awsxray(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.awsxrayexporter.Config"},
			"^awss3(\\/.+)?$":                   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.awss3exporter.Config"},
			"^azuredataexplorer(\\/.+)?$":       {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.azuredataexplorerexporter.Config"},
			"^azuremonitor(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.azuremonitorexporter.Config"},
			"^carbon(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.carbonexporter.Config"},
			"^cassandra(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.cassandraexporter.Config"},
			"^clickhouse(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.clickhouseexporter.Config"},
			"^coralogix(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.coralogixexporter.Config"},
			"^datadog(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.datadogexporter.Config"},
			"^dataset(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.datasetexporter.Config"},
			"^dynatrace(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.dynatraceexporter.config.Config"},
			"^elasticsearch(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.elasticsearchexporter.Config"},
			"^f5cloud(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.f5cloudexporter.Config"},
			"^file(\\/.+)?$":                    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.fileexporter.Config"},
			"^googlecloud(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.googlecloudexporter.Config"},
			"^googlecloudpubsub(\\/.+)?$":       {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.googlecloudpubsubexporter.Config"},
			"^googlemanagedprometheus(\\/.+)?$": {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.googlemanagedprometheusexporter.Config"},
			"^influxdb(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.influxdbexporter.Config"},
			"^instana(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.instanaexporter.Config"},
			"^jaeger(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.jaegerexporter.Config"},
			"^jaegerthrift(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.jaegerthrifthttpexporter.Config"},
			"^kafka(\\/.+)?$":                   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.kafkaexporter.Config"},
			"^loadbalancing(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.loadbalancingexporter.Config"},
			"^logicmonitor(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.logicmonitorexporter.Config"},
			"^logzio(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.logzioexporter.Config"},
			"^loki(\\/.+)?$":                    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.lokiexporter.Config"},
			"^mezmo(\\/.+)?$":                   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.mezmoexporter.Config"},
			"^opencensus(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.opencensusexporter.Config"},
			"^parquet(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.parquetexporter.Config"},
			"^prometheus(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.prometheusexporter.Config"},
			"^prometheusremotewrite(\\/.+)?$":   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.prometheusremotewriteexporter.Config"},
			"^pulsar(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.pulsarexporter.Config"},
			"^sapm(\\/.+)?$":                    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.sapmexporter.Config"},
			"^sentry(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.sentryexporter.Config"},
			"^signalfx(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.signalfxexporter.Config"},
			"^skywalking(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.skywalkingexporter.Config"},
			"^splunkhec(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.splunkhecexporter.Config"},
			"^sumologic(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.sumologicexporter.Config"},
			"^tanzuobservability(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.tanzuobservabilityexporter.Config"},
			"^tencentcloudlogservice(\\/.+)?$":  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.tencentcloudlogserviceexporter.Config"},
			"^zipkin(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.exporter.zipkinexporter.Config"},
		},
	}
	return exporters, &exportersTopLevel
}
