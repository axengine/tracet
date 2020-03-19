package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"time"
	"tracet/gin2micro"
	"tracet/proto/user"
)

func main() {
	cfg := jaegercfg.Configuration{
		ServiceName: "MicroTestService", //自定义服务名称
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "192.168.8.145:5775", //jaeger agent
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	etcdEndpoints := []string{"http://192.168.8.101:2379", "http://192.168.8.121:2379", "http://192.168.8.141:2379"}

	var service web.Service
	service = web.NewService(
		web.Name("tracet.api"),
		web.Address(":12345"),
		web.Registry(etcdv3.NewRegistry(registry.Addrs(etcdEndpoints...))),
	)

	if err := service.Init(); err != nil {
		panic("service.Init:" + err.Error())
	}
	router := gin.Default()
	router.Use(gin.Recovery())
	router.HEAD("/", func(context *gin.Context) {
		context.Status(200)
		return
	})

	v1 := router.Group("/v1")
	v1.Use(gin2micro.TracerWrapper)
	v1.GET("/user", GetUserHandle)

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		panic("service.Run:" + err.Error())
	}

}

func GetUserHandle(c *gin.Context) {
	cli := user.NewUserService("tracet.user", client.DefaultClient)

	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		panic("get context err")
	}

	resp, err := cli.GetUser(ctx, &user.GetUserRequest{
		Id: 1,
	})
	if err != nil {
		c.String(400, "failed = %s", err.Error())
	}

	c.JSON(200, resp)
}
