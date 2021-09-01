package main

import (
	"BookingsApp/helpers"
	"BookingsApp/internal/config"
	"BookingsApp/internal/driver"
	"BookingsApp/internal/handlers"
	"BookingsApp/internal/models"
	"BookingsApp/internal/render"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

const portNumber = ":8083"

// main is the main function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
func run() (*driver.DB, error) {
	//what am I going to
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
	//change this to true when in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session
	//connect to database
	log.Println("Connecting to Database......")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=1234")
	if err != nil {
		log.Fatal("Cannot connect to Database")
	}

	app.UseCache = false
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return nil, err
	}
	log.Println("Connected to Database!!!")
	app.TemplateCache = tc
	//creating repo in main and passing the wide app to handlers
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	return db, nil
}
