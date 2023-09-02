module github.com/ozline/tiktok

go 1.20

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.2.8+incompatible
	github.com/apache/thrift v0.13.0
	github.com/bytedance/sonic v1.8.8
	github.com/cloudwego/fastpb v0.0.4
	github.com/cloudwego/hertz v0.6.7
	github.com/cloudwego/kitex v0.6.2
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/hertz-contrib/opensergo v0.0.1
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20230819133448-76093321aa8e
	github.com/kitex-contrib/registry-etcd v0.1.0
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rabbitmq/amqp091-go v1.8.1
	github.com/redis/go-redis/v9 v9.0.5
	github.com/sirupsen/logrus v1.9.2
	github.com/spf13/viper v1.16.0
	github.com/tinylib/msgp v1.1.8
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	google.golang.org/protobuf v1.31.0
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.3
	gorm.io/hints v1.1.2
	gorm.io/plugin/opentracing v0.0.0-20211220013347-7d2b2af23560
)

require (
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
)

require (
	bou.ke/monkey v1.0.2
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/alibaba/sentinel-golang v1.0.4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.0.0-20230531144706-a12972768317 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/chenzhuoyu/iasm v0.0.0-20230222070914-0b1b64b0e762 // indirect
	github.com/choleraehyq/pid v0.0.17 // indirect
	github.com/cloudwego/configmanager v0.2.0 // indirect
	github.com/cloudwego/frugal v0.1.6 // indirect
	github.com/cloudwego/netpoll v0.4.1 // indirect
	github.com/cloudwego/thriftgo v0.2.11 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/pprof v0.0.0-20220608213341-c488b8fa1db3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/oleiade/lane v1.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.11.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/shirou/gopsutil/v3 v3.21.6 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	github.com/tklauser/numcpus v0.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.etcd.io/etcd/api/v3 v3.5.9 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.9 // indirect
	go.etcd.io/etcd/client/v3 v3.5.9 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/arch v0.2.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.1.0
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/time v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
