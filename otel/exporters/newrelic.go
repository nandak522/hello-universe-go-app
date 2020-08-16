import (
	"log"
	"os"

	"github.com/newrelic/opentelemetry-exporter-go/newrelic"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() {
	exporter, err := newrelic.NewExporter(os.Getenv("APM_APP_NAME"), os.Getenv("APM_LICENSE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	tp, err := trace.NewProvider(trace.WithSyncer(exporter))
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}
