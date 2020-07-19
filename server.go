package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// StartTime gives the start time of server
var StartTime = time.Now()

const defaultAppPort string = "1323"

func uptime() string {
	elapsedTime := time.Since(StartTime)
	return fmt.Sprintf("%d:%d:%d", int(math.Round(elapsedTime.Hours())), int(math.Round(elapsedTime.Minutes())), int(math.Round(elapsedTime.Seconds())))
}

func homePage(c echo.Context) error {
	host, _ := os.Hostname()
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"host":           fmt.Sprintf("[HOST: %s] (uptime: %s)]", host, uptime()),
		"requestHeaders": c.Request().Header,
		"response":       "Hello Universe",
	})
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.HideBanner = true
	e.Debug = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","id":"${id}","host":"${host}",` +
			`,"uri":"${uri}","status":${status},"error":"${error}","latency":${latency_human},` +
			`"bytes_out":${bytes_out}}` + "\n",
		Output: os.Stdout,
	}))
	var port, isEnvVarSet = os.LookupEnv("APP_PORT")
	if !isEnvVarSet {
		port = defaultAppPort
		e.Logger.Info("Port is defaulted to %s", port)
	}
	var enableMonitoring = os.Getenv("ENABLE_MONITORING")
	if enableMonitoring == "" {
		e.Logger.Info("Monitoring is disabled")
	} else {
		var monitoringAppName = os.Getenv("MONITORING_APP_NAME")
		if monitoringAppName == "" {
			panic("MONITORING_APP_NAME env needs to be set, to enable monitoring")
		}
		var monitoringAgentLicenseKey = os.Getenv("MONITORING_LICENSE_KEY")
		if monitoringAgentLicenseKey == "" {
			panic("MONITORING_LICENSE_KEY env needs to be set, to enable monitoring")
		}

		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName(monitoringAppName),
			newrelic.ConfigLicense(monitoringAgentLicenseKey),
		)
		if err != nil {
			panic(err)
		}
		e.Use(nrecho.Middleware(app))
	}
	e.Renderer = renderer
	e.GET("/", homePage)
	e.Logger.Fatal(e.Start(fmt.Sprintf("[::]:%s", port)))
}
