module github.com/pigfall/curdboy_playground

go 1.16

require (
	entgo.io/contrib v0.2.0
	entgo.io/ent v0.9.2-0.20210821141344-368a8f7a2e9a
	github.com/go-kratos/kratos/v2 v2.1.4
	github.com/google/wire v0.5.0
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

replace entgo.io/contrib => github.com/pigfall/contrib v0.1.13
