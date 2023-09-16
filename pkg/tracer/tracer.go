package tracer

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/ozline/tiktok/config"
	"github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string) {
	cfg := &jaegerconfig.Configuration{
		Disabled: false,
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: config.Jaeger.Addr,
		},
	}

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
