package main

import (
	"text/template"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// type (
// 	customContext struct {
// 		echo.Context
// 		serviceDependencyURL *string
// 	}
// )
