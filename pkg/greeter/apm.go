package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func enableAPM(e *echo.Echo) {
	var apmAppName = os.Getenv("APM_APP_NAME")
	if apmAppName == "" {
		e.Logger.Fatal("APM_APP_NAME env needs to be set, to enable application performance monitoring (apm)")
	} else {
		e.Logger.Infof("Will post all transactions data to %s APM_APP_NAME", apmAppName)
	}
	var apmLicenseKey = os.Getenv("APM_LICENSE_KEY")
	if apmLicenseKey == "" {
		e.Logger.Fatal("APM_LICENSE_KEY env needs to be set, to enable application performance monitoring (apm)")
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(apmAppName),
		newrelic.ConfigLicense(apmLicenseKey),
	)
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(nrecho.Middleware(app))
}
