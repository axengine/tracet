package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"tracet/proto/wallet"

	"time"
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
		micro.Name("tracet.wallet"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*3),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
		micro.Registry(etcdv3.NewRegistry(registry.Addrs(etcdEndpoints...))),
	)
	service.Init()

	//Register Service
	_ = wallet.RegisterWalletHandler(service.Server(), NewWalletService())

	if err := service.Run(); err != nil {
		panic("Service.Run()" + err.Error())
		return
	}
}

type WalletService struct {
}

func NewWalletService() *WalletService {
	return &WalletService{}
}

func (s *WalletService) GetWallet(ctx context.Context, in *wallet.GetWalletRequest, out *wallet.GetWalletResponse) error {
	out.Id = in.Id
	out.Symbol = "OLO"
	out.Balance = 1234.5678
	time.Sleep(time.Millisecond * 100)
	return nil
}
