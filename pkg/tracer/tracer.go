package tracer

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string) {
	cfg, _ := jaegerconfig.FromEnv()
	cfg.ServiceName = service

	tracer, _, err := cfg.NewTracer(
		jaegerconfig.Logger(jaeger.StdLogger),
		jaegerconfig.ZipkinSharedRPCSpan(true),
	)

	if err != nil {
		panic(fmt.Sprintf("cannot init jaeger: %v\n", err))
	}

	opentracing.SetGlobalTracer(tracer)
}
