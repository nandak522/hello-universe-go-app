package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func enableAppPerfMonitoring(e *echo.Echo) {
	var monitoringAppName = os.Getenv("MONITORING_APP_NAME")
	if monitoringAppName == "" {
		e.Logger.Fatal("MONITORING_APP_NAME env needs to be set, to enable monitoring")
	}
	var monitoringAgentLicenseKey = os.Getenv("MONITORING_LICENSE_KEY")
	if monitoringAgentLicenseKey == "" {
		e.Logger.Fatal("MONITORING_LICENSE_KEY env needs to be set, to enable monitoring")
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(monitoringAppName),
		newrelic.ConfigLicense(monitoringAgentLicenseKey),
	)
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(nrecho.Middleware(app))
}
