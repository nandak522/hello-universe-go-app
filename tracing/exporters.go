// import (
// 	"log"
// 	"os"

// 	"github.com/newrelic/opentelemetry-exporter-go/newrelic"
// 	"go.opentelemetry.io/otel/api/global"
// 	"go.opentelemetry.io/otel/sdk/trace"
// )

// func initNewrelicExporter(e *echo.Echo, serviceDependencyURL string) {
// 	e.Logger.Info("serviceDependencyURL: ", serviceDependencyURL, " supplied. Hence tracing is enabled")
// 	exporter, err := newrelic.NewExporter("My Service", os.Getenv("NEW_RELIC_API_KEY"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tp, err := trace.NewProvider(trace.WithSyncer(exporter))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	global.SetTraceProvider(tp)
// }
