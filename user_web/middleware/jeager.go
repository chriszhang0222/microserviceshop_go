package middleware


import (
	"github.com/gin-gonic/gin"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"mxshop/user_web/global"
)

func Trace() gin.HandlerFunc{
	return func(context *gin.Context) {
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type: jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans: true,
				LocalAgentHostPort: global.ServerConfig.JaegerInfo.HostPort,
			},
			ServiceName: global.ServerConfig.JaegerInfo.ServiceName,
		}
		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		if err != nil {
			panic(err)
		}
		defer closer.Close()
		startSpan := tracer.StartSpan(context.Request.URL.Path)
		defer startSpan.Finish()
		context.Set("tracer", tracer)
		context.Set("parentSpan", startSpan)
		context.Next()

	}
}