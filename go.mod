module github.com/Gitforxuyang/walleManage

go 1.12

require (
	cloud.google.com/go/bigquery v1.8.0 // indirect
	cloud.google.com/go/pubsub v1.3.1 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/fatih/structs v1.1.0
	github.com/getsentry/sentry-go v0.7.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-errors/errors v1.0.1
	github.com/go-redis/redis/v8 v8.0.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/uuid v1.1.2
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
	google.golang.org/genproto v0.0.0-20200911024640-645f7a48b24f // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
