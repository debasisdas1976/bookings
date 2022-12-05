package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/debasisdas1976/bookings/pkg/config"
	"github.com/debasisdas1976/bookings/pkg/handlers"
	"github.com/debasisdas1976/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production

	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	appConfig.Session = session

	fmt.Println("Main getting invoked")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println(err)
	}

	appConfig.TemplateCache = tc
	appConfig.USeCache = false

	repo := handlers.NewRepo(&appConfig)

	handlers.NewHandlers(repo)

	render.NewTemplates(&appConfig)
	/*
		http.HandleFunc("/", handlers.Repo.Home)
		http.HandleFunc("/about", handlers.Repo.About)
		http.HandleFunc("/contactus", handlers.Repo.ContactUs)
		http.HandleFunc("/orphan", handlers.Repo.Orphan)
		fmt.Println(fmt.Sprintf("Starting Application Server at %s", portNumber))

		_ = http.ListenAndServe(portNumber, nil)
	*/

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
