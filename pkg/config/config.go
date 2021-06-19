package config

import "html/template"

//Appcongi is configuration stuct for the app
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
