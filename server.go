package main

import (
	"embed"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

// Embed file content as a string
//go:embed static/*
var staticContent embed.FS

// Embed templates
//go:embed templates
var templatesContent embed.FS

// StartTime gives the start time of server
var StartTime = time.Now()

const defaultAppPort string = "1323"

func getLogLevel(suppliedLogLevel string) log.Level {
	if suppliedLogLevel == "" {
		return log.InfoLevel
	}
	switch suppliedLogLevel {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	default:
		log.Fatal("Please supply a valid log level")
	}
	return log.DebugLevel
}

func uptime() string {
	elapsedTime := time.Since(StartTime)
	return fmt.Sprintf("%d:%d:%d", int(math.Round(elapsedTime.Hours())), int(math.Round(elapsedTime.Minutes())), int(math.Round(elapsedTime.Seconds())))
}

func homePageHandler(rw http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path
	log.Debug("Serving request for path: ", path)
	rw.Header().Add("Content-type", "text/html")

	homePageTemplate, err := template.ParseFS(templatesContent, "templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	host, _ := os.Hostname()
	homePageTemplate.Execute(rw, map[string]interface{}{
		"uptime":         uptime(),
		"host":           host,
		"requestHeaders": r.Header,
		"response":       "Hello Universe",
	})
}

func main() {
	log.SetFormatter(&log.JSONFormatter{
		DisableHTMLEscape: true,
		TimestampFormat:   time.RFC3339,
	})

	var requiredLogLevel string
	flag.StringVarP(&requiredLogLevel, "log-level", "l", "info", "Required log level: debug/info/warn/error. Defaults to info")
	var printHelp bool
	flag.BoolVarP(&printHelp, "help", "h", false, "Prints this help content.")
	var printVersion bool
	flag.BoolVarP(&printVersion, "version", "v", false, "Prints the version of Heva.")

	flag.Parse()
	if printHelp {
		flag.Usage()
		return
	}
	if printVersion {
		fmt.Println("v" + strings.Join(VERSION[:], "."))
		os.Exit(0)
	}
	log.SetLevel(getLogLevel(requiredLogLevel))
	log.SetOutput(os.Stdout)

	var staticFS = http.FS(staticContent)
	fs := http.FileServer(staticFS)
	http.Handle("/static/", fs)
	http.HandleFunc("/", homePageHandler)

	var port, isEnvVarSet = os.LookupEnv("APP_PORT")
	if !isEnvVarSet {
		port = defaultAppPort
	}
	log.Info("http server is ready to serve at port ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
