package main

import (
	"embed"
	"encoding/json"
	"fmt"
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

type HomePageResponse struct {
	RequestHeaders map[string][]string `json:"requestHeaders"`
	Content        string              `json:"content"`
	Uptime         string              `json:"uptime"`
	Host           string              `json:"host"`
}

func homePageHandler(rw http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path
	log.Debug("Serving request for path: ", path)
	host, _ := os.Hostname()
	content := "Hello Universe"
	uptime := time.Now().Sub(StartTime).Round(time.Second)
	response := HomePageResponse{
		RequestHeaders: r.Header,
		Content:        content,
		Uptime:         uptime.String(),
		Host:           host,
	}

	if strings.Contains(r.Header.Get("Content-Type"), "json") {
		rw.Header().Add("Content-type", "application/json")
		// rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(response)
	} else {
		rw.Header().Add("Content-type", "text/html")
		homePageTemplate, err := template.ParseFS(templatesContent, "templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		homePageTemplate.Execute(rw, map[string]interface{}{
			"uptime":         response.Uptime,
			"host":           response.Host,
			"requestHeaders": response.RequestHeaders,
			"content":        response.Content,
		})
	}
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
	flag.BoolVarP(&printVersion, "version", "v", false, "Prints the version of hello-universe-app.")

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
