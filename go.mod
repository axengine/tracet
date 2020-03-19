module tracet

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-log/log v0.1.0
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro v1.13.1
	github.com/micro/go-plugins v1.3.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/zap v1.14.1
)

replace (
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/google/pprof v0.0.0-20190515194954-54271f7e092f
	golang.org/x/image v0.0.0-00010101000000-000000000000 => github.com/golang/image v0.0.0-20180708004352-c73c2afc3b81
)
