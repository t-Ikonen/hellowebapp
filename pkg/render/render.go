package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/t-Ikonen/hellowebapp/pkg/config"
	"github.com/t-Ikonen/hellowebapp/pkg/models"
)

var functions = template.FuncMap{}

var appConfig *config.AppConfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

//NewTemplates sets the package for the template package
func NewTemplates(a *config.AppConfig) {
	appConfig = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, tmplD *models.TemplateData) {
	var tmplCache map[string]*template.Template

	if appConfig.UseCache {
		tmplCache = appConfig.TemplateCache
	} else {
		tmplCache, _ = CreateTemplateCache()
	}

	t, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	tmplD = AddDefaultData(tmplD)

	_ = t.Execute(buf, tmplD)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

//CreateTemplateCache creates template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	println("pages: ", pages)
	for _, page := range pages {
		name := filepath.Base(page)
		//fmt.Println("page filelistassa on", page, "ja name on ", name)

		tmplSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("Template setin luonti")
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			fmt.Println("Matches")
			return myCache, err
		}
		if len(matches) > 0 {
			tmplSet, err = tmplSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {

				return myCache, err
			}
		}
		myCache[name] = tmplSet
	}
	return myCache, nil
}
