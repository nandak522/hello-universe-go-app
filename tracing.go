package main

import "github.com/labstack/echo/v4"

func initTracer(e *echo.Echo, serviceDependencyURL string) {
	e.Logger.Info("serviceDependencyURL: ", serviceDependencyURL, " supplied. Hence tracing is enabled")
}
