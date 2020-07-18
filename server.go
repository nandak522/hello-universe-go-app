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
)

// StartTime gives the start time of server
var StartTime = time.Now()

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
	e.Renderer = renderer
	e.GET("/", homePage)
	e.Logger.Fatal(e.Start("[::]:1323"))
}
