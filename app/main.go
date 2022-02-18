package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	app "github.com/k0pernicus/go-photoaccess/internal"
	annotationHandlers "github.com/k0pernicus/go-photoaccess/internal/handlers/annotation"
	photoHandlers "github.com/k0pernicus/go-photoaccess/internal/handlers/photo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func main() {
	// Read configuration file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Errorf("Cannot read configuration file at root: %s", err.Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &app.C)
	if err != nil {
		log.Errorf("Unmarshal configuration error: %s", err.Error())
		os.Exit(1)
	}

	appCtx := context.Background()

	conn, err := pgxpool.Connect(appCtx, app.C.DB.String())
	if err != nil {
		log.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	app.DB = conn

	// Register handlers
	log.Debug("Registering handlers... ")
	router := mux.NewRouter()

	// Photo handlers
	router.HandleFunc("/photo", photoHandlers.Create).Methods("POST")
	router.HandleFunc("/photo/{id}", photoHandlers.Delete).Methods("DELETE")
	router.HandleFunc("/photo/{id}", photoHandlers.Get).Methods("GET")

	// Annotation handlers
	router.HandleFunc("/photo/{id}/annotation", annotationHandlers.Create).Methods("POST")
	router.HandleFunc("/photo/{id}/annotation/{id}", annotationHandlers.Delete).Methods("DELETE")
	router.HandleFunc("/photo/{id}/annotation/{id}", annotationHandlers.Get).Methods("GET")

	appAddr := app.C.App.String()

	// Create server
	srv := &http.Server{
		Handler:      router,
		Addr:         appAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server is running on %s\n", appAddr)

	errorSignal := make(chan bool, 1)

	// Let the service crash if the connection to DB fails
	go func() {
		for {
			if err := app.DB.Ping(appCtx); err != nil {
				log.Errorf("Connection to DB dropped (%+v)", err)
				errorSignal <- true
				return
			}
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		// Serve
		if err := srv.ListenAndServe(); err != nil {
			log.Errorf("Error running the service: %+v\n", err)
			errorSignal <- true
		}
	}()

	// Wait until error
	<-errorSignal

	if err := srv.Shutdown(appCtx); err != nil {
		log.Errorf("Error when shutdown the service: %+v\n", err)
	}
}
