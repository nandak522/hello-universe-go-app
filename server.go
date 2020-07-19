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
	log "github.com/labstack/gommon/log"
	flag "github.com/spf13/pflag"
)

// StartTime gives the start time of server
var StartTime = time.Now()

const defaultAppPort string = "1323"

func uptime() string {
	elapsedTime := time.Since(StartTime)
	return fmt.Sprintf("%d:%d:%d", int(math.Round(elapsedTime.Hours())), int(math.Round(elapsedTime.Minutes())), int(math.Round(elapsedTime.Seconds())))
}

func homePage(serviceDependencyURL *string) echo.HandlerFunc {
	return func(c echo.Context) error {
		host, _ := os.Hostname()
		if *serviceDependencyURL != "" {
			response, err := http.Get(*serviceDependencyURL)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
		}
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"host":           fmt.Sprintf("[HOST: %s] (uptime: %s)]", host, uptime()),
			"requestHeaders": c.Request().Header,
			"response":       "Hello Universe.",
		})
	}
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	var serviceDependencyURL string
	flag.StringVarP(&serviceDependencyURL, "service-dep-url", "s", "", "External Service Dependency Url")
	flag.Parse()
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.HideBanner = true
	e.Debug = true
	if serviceDependencyURL == "" {
		e.Logger.Warn("serviceDependencyURL not supplied. Hence tracing is disabled")
	} else {
		// e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		// 	return func(c echo.Context) error {
		// 		cc := &customContext{c, &serviceDependencyURL}
		// 		return h(cc)
		// 	}
		// })
		initTracer(e, serviceDependencyURL)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","id":"${id}","host":"${host}",` +
			`,"uri":"${uri}","status":${status},"error":"${error}","latency":${latency_human},` +
			`"bytes_out":${bytes_out}}` + "\n",
		Output: os.Stdout,
	}))
	e.Logger.SetLevel(log.DEBUG)
	var port, isEnvVarSet = os.LookupEnv("APP_PORT")
	if !isEnvVarSet {
		port = defaultAppPort
		e.Logger.Infof("Port is defaulted to %s", port)
	}
	var enableAppPerfMonitoring = os.Getenv("ENABLE_APM")
	if enableAppPerfMonitoring == "" {
		e.Logger.Warn("Application Performance Monitoring is disabled")
	} else {
		enableAPM(e)
	}

	e.Renderer = renderer
	e.GET("/", homePage(&serviceDependencyURL))
	e.Logger.Fatal(e.Start(fmt.Sprintf("[::]:%s", port)))
}
