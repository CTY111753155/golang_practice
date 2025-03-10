package config

import (
	"html/template"
	"log"
	"myapp/internal/models"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Inproduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
