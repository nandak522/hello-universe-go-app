package main

import (
	log "github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/exporters/trace/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initOpenTelemetry(e *echo.Echo, serviceDependencyURL string) {
	e.Logger.Info("serviceDependencyURL: ", serviceDependencyURL, " supplied. Hence tracing is enabled")
	exporter, err := stdout.NewExporter(stdout.Options{PrettyPrint: true})
	if err != nil {
		log.Fatal(err)
	}
	tp, err := sdktrace.NewProvider(sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(exporter))
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}
