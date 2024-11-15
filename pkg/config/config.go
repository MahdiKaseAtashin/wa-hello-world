package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Port          string
	InfoLog       *log.Logger
}
