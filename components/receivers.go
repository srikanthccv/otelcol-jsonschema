package components

import (
	"github.com/knadh/koanf/v2"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/bigipreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudflarereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dotnetdiagnosticsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filestatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/haproxyreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otlpjsonfilereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusexecreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snowflakereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zookeeperreceiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"

	"github.com/srikanthccv/otel/collector/schema/jsonschema"
)

func Receivers(r *jsonschema.Reflector) ([]*koanf.Koanf, *jsonschema.Schema) {
	activedirectoryds := KoanfLoad(r.Reflect(&activedirectorydsreceiver.Config{}).MarshalJSON())
	aerospikereceiver := KoanfLoad(r.Reflect(&aerospikereceiver.Config{}).MarshalJSON())
	apache := KoanfLoad(r.Reflect(&apachereceiver.Config{}).MarshalJSON())
	apachespark := KoanfLoad(r.Reflect(&apachesparkreceiver.Config{}).MarshalJSON())
	awscloudwatch := KoanfLoad(r.Reflect(&awscloudwatchreceiver.Config{}).MarshalJSON())
	awscontainerinsight := KoanfLoad(r.Reflect(&awscontainerinsightreceiver.Config{}).MarshalJSON())
	awsecscontainermetrics := KoanfLoad(r.Reflect(&awsecscontainermetricsreceiver.Config{}).MarshalJSON())
	awsfirehose := KoanfLoad(r.Reflect(&awsfirehosereceiver.Config{}).MarshalJSON())
	awsxray := KoanfLoad(r.Reflect(&awsxrayreceiver.Config{}).MarshalJSON())
	azureblob := KoanfLoad(r.Reflect(&azureblobreceiver.Config{}).MarshalJSON())
	azureeventhub := KoanfLoad(r.Reflect(&azureeventhubreceiver.Config{}).MarshalJSON())
	azuremonitor := KoanfLoad(r.Reflect(&azuremonitorreceiver.Config{}).MarshalJSON())
	bigip := KoanfLoad(r.Reflect(&bigipreceiver.Config{}).MarshalJSON())
	carbon := KoanfLoad(r.Reflect(&carbonreceiver.Config{}).MarshalJSON())
	chrony := KoanfLoad(r.Reflect(&chronyreceiver.Config{}).MarshalJSON())
	cloudflare := KoanfLoad(r.Reflect(&cloudflarereceiver.Config{}).MarshalJSON())
	cloudfoundry := KoanfLoad(r.Reflect(&cloudfoundryreceiver.Config{}).MarshalJSON())
	collectd := KoanfLoad(r.Reflect(&collectdreceiver.Config{}).MarshalJSON())
	couchdb := KoanfLoad(r.Reflect(&couchdbreceiver.Config{}).MarshalJSON())
	datadog := KoanfLoad(r.Reflect(&datadogreceiver.Config{}).MarshalJSON())
	dockerstats := KoanfLoad(r.Reflect(&dockerstatsreceiver.Config{}).MarshalJSON())
	dotnetdiagnostics := KoanfLoad(r.Reflect(&dotnetdiagnosticsreceiver.Config{}).MarshalJSON())
	elasticsearch := KoanfLoad(r.Reflect(&elasticsearchreceiver.Config{}).MarshalJSON())
	expvar := KoanfLoad(r.Reflect(&expvarreceiver.Config{}).MarshalJSON())
	filelog := KoanfLoad(r.Reflect(&filelogreceiver.FileLogConfig{}).MarshalJSON())
	filestats := KoanfLoad(r.Reflect(&filestatsreceiver.Config{}).MarshalJSON())
	file := KoanfLoad(r.Reflect(&filereceiver.Config{}).MarshalJSON())
	flinkmetrics := KoanfLoad(r.Reflect(&flinkmetricsreceiver.Config{}).MarshalJSON())
	fluentforward := KoanfLoad(r.Reflect(&fluentforwardreceiver.Config{}).MarshalJSON())
	googlecloudpubsub := KoanfLoad(r.Reflect(&googlecloudpubsubreceiver.Config{}).MarshalJSON())
	googlecloudspanner := KoanfLoad(r.Reflect(&googlecloudspannerreceiver.Config{}).MarshalJSON())
	haproxy := KoanfLoad(r.Reflect(&haproxyreceiver.Config{}).MarshalJSON())
	hostmetrics := KoanfLoad(r.Reflect(&hostmetricsreceiver.Config{}).MarshalJSON())
	httpcheck := KoanfLoad(r.Reflect(&httpcheckreceiver.Config{}).MarshalJSON())
	influxdb := KoanfLoad(r.Reflect(&influxdbreceiver.Config{}).MarshalJSON())
	iis := KoanfLoad(r.Reflect(&iisreceiver.Config{}).MarshalJSON())
	jaeger := KoanfLoad(r.Reflect(&jaegerreceiver.Config{}).MarshalJSON())
	jmx := KoanfLoad(r.Reflect(&jmxreceiver.Config{}).MarshalJSON())
	journald := KoanfLoad(r.Reflect(&journaldreceiver.JournaldConfig{}).MarshalJSON())
	k8scluster := KoanfLoad(r.Reflect(&k8sclusterreceiver.Config{}).MarshalJSON())
	k8sevents := KoanfLoad(r.Reflect(&k8seventsreceiver.Config{}).MarshalJSON())
	k8sobjects := KoanfLoad(r.Reflect(&k8sobjectsreceiver.Config{}).MarshalJSON())
	kafkametrics := KoanfLoad(r.Reflect(&kafkametricsreceiver.Config{}).MarshalJSON())
	kafka := KoanfLoad(r.Reflect(&kafkareceiver.Config{}).MarshalJSON())
	kubeletstats := KoanfLoad(r.Reflect(&kubeletstatsreceiver.Config{}).MarshalJSON())
	loki := KoanfLoad(r.Reflect(&lokireceiver.Config{}).MarshalJSON())
	memcached := KoanfLoad(r.Reflect(&memcachedreceiver.Config{}).MarshalJSON())
	mongodbatlas := KoanfLoad(r.Reflect(&mongodbatlasreceiver.Config{}).MarshalJSON())
	mongodb := KoanfLoad(r.Reflect(&mongodbreceiver.Config{}).MarshalJSON())
	mysql := KoanfLoad(r.Reflect(&mysqlreceiver.Config{}).MarshalJSON())
	nginx := KoanfLoad(r.Reflect(&nginxreceiver.Config{}).MarshalJSON())
	nsxt := KoanfLoad(r.Reflect(&nsxtreceiver.Config{}).MarshalJSON())
	opencensus := KoanfLoad(r.Reflect(&opencensusreceiver.Config{}).MarshalJSON())
	oracledb := KoanfLoad(r.Reflect(&oracledbreceiver.Config{}).MarshalJSON())
	otlp := KoanfLoad(r.Reflect(&otlpreceiver.Config{}).MarshalJSON())
	otlpjsonfile := KoanfLoad(r.Reflect(&otlpjsonfilereceiver.Config{}).MarshalJSON())
	podman := KoanfLoad(r.Reflect(&podmanreceiver.Config{}).MarshalJSON())
	postgresql := KoanfLoad(r.Reflect(&postgresqlreceiver.Config{}).MarshalJSON())
	prometheusexec := KoanfLoad(r.Reflect(&prometheusexecreceiver.Config{}).MarshalJSON())
	prometheus := KoanfLoad(r.Reflect(&prometheusreceiver.Config{}).MarshalJSON())
	purefa := KoanfLoad(r.Reflect(&purefareceiver.Config{}).MarshalJSON())
	purefb := KoanfLoad(r.Reflect(&purefbreceiver.Config{}).MarshalJSON())
	pulsar := KoanfLoad(r.Reflect(&pulsarreceiver.Config{}).MarshalJSON())
	rabbitmq := KoanfLoad(r.Reflect(&rabbitmqreceiver.Config{}).MarshalJSON())
	receivercreator := KoanfLoad(r.Reflect(&receivercreator.Config{}).MarshalJSON())
	redis := KoanfLoad(r.Reflect(&redisreceiver.Config{}).MarshalJSON())
	riak := KoanfLoad(r.Reflect(&riakreceiver.Config{}).MarshalJSON())
	sapm := KoanfLoad(r.Reflect(&sapmreceiver.Config{}).MarshalJSON())
	signalfx := KoanfLoad(r.Reflect(&signalfxreceiver.Config{}).MarshalJSON())
	simpleprometheus := KoanfLoad(r.Reflect(&simpleprometheusreceiver.Config{}).MarshalJSON())
	skywalking := KoanfLoad(r.Reflect(&skywalkingreceiver.Config{}).MarshalJSON())
	snowflake := KoanfLoad(r.Reflect(&snowflakereceiver.Config{}).MarshalJSON())
	solace := KoanfLoad(r.Reflect(&solacereceiver.Config{}).MarshalJSON())
	splunkhec := KoanfLoad(r.Reflect(&splunkhecreceiver.Config{}).MarshalJSON())
	sqlquery := KoanfLoad(r.Reflect(&sqlqueryreceiver.Config{}).MarshalJSON())
	sqlserver := KoanfLoad(r.Reflect(&sqlserverreceiver.Config{}).MarshalJSON())
	sshcheck := KoanfLoad(r.Reflect(&sshcheckreceiver.Config{}).MarshalJSON())
	statsd := KoanfLoad(r.Reflect(&statsdreceiver.Config{}).MarshalJSON())
	// syslog := KoanfLoad(r.Reflect(&syslogreceiver.Config{}).MarshalJSON())
	// tcplog := KoanfLoad(r.Reflect(&tcplogreceiver.Config{}).MarshalJSON())
	// udplog := KoanfLoad(r.Reflect(&udplogreceiver.Config{}).MarshalJSON())
	vcenter := KoanfLoad(r.Reflect(&vcenterreceiver.Config{}).MarshalJSON())
	wavefront := KoanfLoad(r.Reflect(&wavefrontreceiver.Config{}).MarshalJSON())
	snmp := KoanfLoad(r.Reflect(&snmpreceiver.Config{}).MarshalJSON())
	windowseventlog := KoanfLoad(r.Reflect(&windowseventlogreceiver.WindowsLogConfig{}).MarshalJSON())
	windowsperfcounters := KoanfLoad(r.Reflect(&windowsperfcountersreceiver.Config{}).MarshalJSON())
	zipkin := KoanfLoad(r.Reflect(&zipkinreceiver.Config{}).MarshalJSON())
	zookeeper := KoanfLoad(r.Reflect(&zookeeperreceiver.Config{}).MarshalJSON())

	receivers := []*koanf.Koanf{
		activedirectoryds,
		aerospikereceiver,
		apache,
		apachespark,
		awscloudwatch,
		awscontainerinsight,
		awsecscontainermetrics,
		awsfirehose,
		awsxray,
		azureblob,
		azureeventhub,
		azuremonitor,
		bigip,
		carbon,
		chrony,
		cloudflare,
		cloudfoundry,
		collectd,
		couchdb,
		datadog,
		dockerstats,
		dotnetdiagnostics,
		elasticsearch,
		expvar,
		filelog,
		filestats,
		file,
		flinkmetrics,
		fluentforward,
		googlecloudpubsub,
		googlecloudspanner,
		haproxy,
		hostmetrics,
		httpcheck,
		influxdb,
		iis,
		jaeger,
		jmx,
		journald,
		k8scluster,
		k8sevents,
		k8sobjects,
		kafkametrics,
		kafka,
		kubeletstats,
		loki,
		memcached,
		mongodbatlas,
		mongodb,
		mysql,
		nginx,
		nsxt,
		opencensus,
		oracledb,
		otlp,
		otlpjsonfile,
		podman,
		postgresql,
		prometheusexec,
		prometheus,
		purefa,
		purefb,
		pulsar,
		rabbitmq,
		receivercreator,
		redis,
		riak,
		sapm,
		signalfx,
		simpleprometheus,
		skywalking,
		snowflake,
		solace,
		splunkhec,
		sqlquery,
		sqlserver,
		sshcheck,
		statsd,
		// syslog,
		// tcplog,
		// udplog,
		vcenter,
		wavefront,
		snmp,
		windowseventlog,
		windowsperfcounters,
		zipkin,
		zookeeper,
	}

	receiversTopLevel := jsonschema.Schema{
		MinProperties: 1,
		Type:          "object",
		PatternProperties: map[string]*jsonschema.Schema{
			"^activedirectoryds(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.activedirectorydsreceiver.Config"},
			"^aerospikereceiver(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.aerospikereceiver.Config"},
			"^apache(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.apachereceiver.Config"},
			"^apachespark(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.apachesparkreceiver.Config"},
			"^awscloudwatch(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awscloudwatchreceiver.Config"},
			"^awscontainerinsight(\\/.+)?$":    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awscontainerinsightreceiver.Config"},
			"^awsecscontainermetrics(\\/.+)?$": {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awsecscontainermetricsreceiver.Config"},
			"^awsfirehose(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awsfirehosereceiver.Config"},
			"^awsxray(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.awsxrayreceiver.Config"},
			"^azureblob(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.azureblobreceiver.Config"},
			"^azureeventhub(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.azureeventhubreceiver.Config"},
			"^azuremonitor(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.azuremonitorreceiver.Config"},
			"^bigip(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.bigipreceiver.Config"},
			"^carbon(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.carbonreceiver.Config"},
			"^chrony(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.chronyreceiver.Config"},
			"^cloudflare(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.cloudflarereceiver.Config"},
			"^cloudfoundry(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.cloudfoundryreceiver.Config"},
			"^collectd(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.collectdreceiver.Config"},
			"^couchdb(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.couchdbreceiver.Config"},
			"^datadog(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.datadogreceiver.Config"},
			"^dockerstats(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.dockerstatsreceiver.Config"},
			"^dotnetdiagnostics(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.dotnetdiagnosticsreceiver.Config"},
			"^elasticsearch(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.elasticsearchreceiver.Config"},
			"^expvar(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.expvarreceiver.Config"},
			"^filelog(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.filelogreceiver.FileLogConfig"},
			"^filestats(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.filestatsreceiver.Config"},
			"^file(\\/.+)?$":                   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.filereceiver.Config"},
			"^flinkmetrics(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.flinkmetricsreceiver.Config"},
			"^fluentforward(\\/.+)?$":          {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.fluentforwardreceiver.Config"},
			"^googlecloudpubsub(\\/.+)?$":      {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.googlecloudpubsubreceiver.Config"},
			"^googlecloudspanner(\\/.+)?$":     {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.googlecloudspannerreceiver.Config"},
			"^haproxy(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.haproxyreceiver.Config"},
			"^hostmetrics(\\/.+)?$":            {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.hostmetricsreceiver.Config"},
			"^httpcheck(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.httpcheckreceiver.Config"},
			"^influxdb(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.influxdbreceiver.Config"},
			"^iis(\\/.+)?$":                    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.iisreceiver.Config"},
			"^jaeger(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.jaegerreceiver.Config"},
			"^jmx(\\/.+)?$":                    {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.jmxreceiver.Config"},
			"^journald(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.journaldreceiver.JournaldConfig"},
			"^k8scluster(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.k8sclusterreceiver.Config"},
			"^k8sevents(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.k8seventsreceiver.Config"},
			"^k8sobjects(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.k8sobjectsreceiver.Config"},
			"^kafkametrics(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.kafkametricsreceiver.Config"},
			"^kafka(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.kafkareceiver.Config"},
			"^kubeletstats(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.kubeletstatsreceiver.Config"},
			"^loki(\\/.+)?$":                   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.lokireceiver.Config"},
			"^memcached(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.memcachedreceiver.Config"},
			"^mongodbatlas(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.mongodbatlasreceiver.Config"},
			"^mongodb(\\/.+)?$":                {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.mongodbreceiver.Config"},
			"^mysql(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.mysqlreceiver.Config"},
			"^nginx(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.nginxreceiver.Config"},
			"^nsxt(\\/.+)?$":                   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.nsxtreceiver.Config"},
			"^opencensus(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.opencensusreceiver.Config"},
			"^oracledb(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.oracledbreceiver.Config"},
			"^otlp(\\/.+)?$":                   {Ref: "#/$defs/go.opentelemetry.io.collector.receiver.otlpreceiver.Config"},
			"^otlpjsonfile(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.otlpjsonfilereceiver.Config"},
			"^podman(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.podmanreceiver.Config"},
			"^postgresql(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.postgresqlreceiver.Config"},
			"^prometheusexec(\\/.+)?$":         {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.prometheusexecreceiver.Config"},
			"^prometheus(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.prometheusreceiver.Config"},
			"^purefa(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.purefareceiver.Config"},
			"^purefb(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.purefbreceiver.Config"},
			"^pulsar(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.pulsarreceiver.Config"},
			"^rabbitmq(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.rabbitmqreceiver.Config"},
			"^receivercreator(\\/.+)?$":        {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.receivercreator.Config"},
			"^redis(\\/.+)?$":                  {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.redisreceiver.Config"},
			"^riak(\\/.+)?$":                   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.riakreceiver.Config"},
			"^sapm(\\/.+)?$":                   {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.sapmreceiver.Config"},
			"^signalfx(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.signalfxreceiver.Config"},
			"^simpleprometheus(\\/.+)?$":       {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.simpleprometheusreceiver.Config"},
			"^skywalking(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.skywalkingreceiver.Config"},
			"^snowflake(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.snowflakereceiver.Config"},
			"^solace(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.solacereceiver.Config"},
			"^splunkhec(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.splunkhecreceiver.Config"},
			"^sqlquery(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.sqlqueryreceiver.Config"},
			"^sqlserver(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.sqlserverreceiver.Config"},
			"^sshcheck(\\/.+)?$":               {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.sshcheckreceiver.Config"},
			"^statsd(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.statsdreceiver.Config"},
			// "^syslog(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.syslogreceiver.Config"},
			// "^tcplog(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.tcplogreceiver.Config"},
			// "^udplog(\\/.+)?$":                 {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.udplogreceiver.Config"},
			"^vcenter(\\/.+)?$":             {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.vcenterreceiver.Config"},
			"^wavefront(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.wavefrontreceiver.Config"},
			"^windowseventlog(\\/.+)?$":     {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.windowseventlogreceiver.WindowsLogConfig"},
			"^windowsperfcounters(\\/.+)?$": {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.windowsperfcountersreceiver.Config"},
			"^zipkin(\\/.+)?$":              {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.zipkinreceiver.Config"},
			"^zookeeper(\\/.+)?$":           {Ref: "#/$defs/github.com.open-telemetry.opentelemetry-collector-contrib.receiver.zookeeperreceiver.Config"},
		},
	}
	return receivers, &receiversTopLevel
}
