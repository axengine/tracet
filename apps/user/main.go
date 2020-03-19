package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"tracet/proto/wallet"

	"time"
	"tracet/proto/user"
)

func main() {
	// jaeger
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

	etcdEndpoints := []string{"http://192.168.8.101:2379", "http://192.168.8.121:2379", "http://192.168.8.141:2379"}

	var service micro.Service
	service = micro.NewService(
		micro.Name("tracet.user"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*3),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
		micro.Registry(etcdv3.NewRegistry(registry.Addrs(etcdEndpoints...))),
	)
	service.Init()

	//Register Service
	_ = user.RegisterUserHandler(service.Server(), NewUserService())

	if err := service.Run(); err != nil {
		panic("Service.Run()" + err.Error())
		return
	}
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context, in *user.GetUserRequest, out *user.GetUserResponse) error {
	out.Id = in.Id
	out.Name = "hehe"
	time.Sleep(time.Millisecond * 100)

	cli := wallet.NewWalletService("tracet.wallet", client.DefaultClient)
	resp, err := cli.GetWallet(ctx, &wallet.GetWalletRequest{
		Id: in.Id,
	})
	if err != nil {
		return err
	}
	out.Symbol = resp.Symbol
	out.Balance = resp.Balance
	return nil
}
