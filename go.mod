module github.com/Martin1992-ss/deMicro

go 1.14

replace (
	github.com/Martin1992-ss/deMicro/common => ./deMicro/common
	github.com/Martin1992-ss/deMicro/config => ./deMicro/config
	github.com/Martin1992-ss/deMicro/consignment/handler => ./deMicro/consignment/handler
	github.com/Martin1992-ss/deMicro/interface-center => ./deMicro/interface-center

)

require (
	github.com/HdrHistogram/hdrhistogram-go v0.9.0 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/armon/go-metrics v0.3.4 // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/consul/api v1.7.0 // indirect
	github.com/hashicorp/go-hclog v0.14.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.9.4 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/lucas-clemente/quic-go v0.18.0 // indirect
	github.com/marten-seemann/qtls-go1-15 v0.1.1 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/broker/kafka v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/tcp v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/breaker/hystrix v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/ratelimiter/uber v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb // indirect
	github.com/micro/protoc-gen-micro/v2 v2.3.0 // indirect
	github.com/miekg/dns v1.1.31 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0 // indirect
	golang.org/x/net v0.0.0-20201002202402-0a1ea396d57c
	golang.org/x/tools v0.0.0-20201002184944-ecd9fd270d5d // indirect
	google.golang.org/genproto v0.0.0-20201002142447-3860012362da // indirect
	google.golang.org/grpc/examples v0.0.0-20201002194053-b2c5f4a808fd // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.3.0
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect

)
