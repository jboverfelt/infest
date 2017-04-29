package main

import (
	"os"

	_ "github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
	_ "github.com/robfig/cron"
)

type props struct {
	ClosuresURL string
	DatabaseURL string
	Port        string
}

func main() {
	//p := fillFromEnv()
}

func fillFromEnv() *props {
	p := &props{
		ClosuresURL: os.Getenv("CLOSURES_URL"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	if p.ClosuresURL == "" {
		panic("Closures URL not defined")
	}

	if p.DatabaseURL == "" {
		panic("Database URL not defined")
	}

	if p.Port == "" {
		panic("Port not defined")
	}

	return p
}
