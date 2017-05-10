package main

import (
	"log"
	"net/http"
	"os"

	"github.com/markbates/pop"
	"github.com/robfig/cron"

	"github.com/gorilla/mux"
)

type props struct {
	ClosuresURL string
	Port        string
	Schedule    string
	GoEnv       string
}

const closureTimeFmt = "1/2/2006"

func checkFatalErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	p := fillFromEnv()

	db, err := pop.Connect(p.GoEnv)
	checkFatalErr(err)

	err = db.MigrateUp("migrations")
	checkFatalErr(err)

	loadFunc := createLoadFunc(db, p)

	// run load process on startup to make sure we're up
	// to date when heroku wakes the app up
	loadFunc()

	startCron(p, loadFunc)

	env := &Env{
		DB: db,
	}

	r := mux.NewRouter()
	r.Handle("/", NewHandler(env, all)).Methods("GET")

	http.ListenAndServe(":"+p.Port, r)
}

func fillFromEnv() *props {
	p := &props{
		ClosuresURL: os.Getenv("CLOSURES_URL"),
		Port:        os.Getenv("PORT"),
		GoEnv:       os.Getenv("GO_ENV"),
		Schedule:    os.Getenv("CLOSURES_SCHEDULE"),
	}

	if p.ClosuresURL == "" {
		panic("Closures URL not defined")
	}

	if p.Port == "" {
		panic("Port not defined")
	}

	if p.GoEnv == "" {
		log.Printf("Env not defined, setting to development")
		p.GoEnv = "development"
	}

	if p.Schedule == "" {
		log.Printf("Schedule not defined, using @every 1m")
		p.Schedule = "@every 1m"
	}

	return p
}

func startCron(p *props, cronFunc func()) {
	c := cron.New()

	err := c.AddFunc(p.Schedule, cronFunc)

	checkFatalErr(err)

	c.Start()
}

func createLoadFunc(db *pop.Connection, p *props) func() {
	return func() {
		log.Printf("Starting load function")
		closures, err := fetchClosures(p.ClosuresURL)

		if err != nil {
			log.Printf("Failed to fetch closures: %v\n", err)
			return
		}

		err = insertIntoDb(closures, db)

		if err != nil {
			log.Printf("Failed to insert into db %v\n", err)
		}
	}
}
