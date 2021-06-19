package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/t-Ikonen/hellowebapp/pkg/config"
	"github.com/t-Ikonen/hellowebapp/pkg/handlers"
	"github.com/t-Ikonen/hellowebapp/pkg/render"
)

const portNum = ":8080"

//Main of HelloWeb app
func main() {

	var appCnf config.AppConfig

	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Printf("Error crating template configuration, error %s \n", err)
		//fmt.Println(fmt.Sprintf("Error crating template configuration, error %s \n", err))
	}
	appCnf.TemplateCache = tmplCache
	appCnf.UseCache = false

	repo := handlers.NewRepo(&appCnf)
	handlers.NewHandlers(repo)

	render.NewTemplates(&appCnf)

	srv := &http.Server{
		Addr:    portNum,
		Handler: Routes(&appCnf),
	}
	fmt.Printf("Starting app on port %s for your pleasure \n", portNum)
	//fmt.Println(fmt.Sprintf("Starting app on port %s for your pleasure \n", portNum))
	err = srv.ListenAndServe()
	log.Fatal(err)

}
