package main

import (
	"github.com/go-chi/chi"
	"github.com/jexroid/gopi/internal/router"
	"net/http"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	var r = chi.NewRouter()
	router.Handler(r)

	log.Info("go server is running on http://localhost:8000")

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Error(err)
	}
}
