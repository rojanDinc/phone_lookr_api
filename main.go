package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"phone_lookr/controller"
	"phone_lookr/controller/middleware"
	"phone_lookr/persistence"
	"phone_lookr/service"
	"phone_lookr/util"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func closeApplicationHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

var environment *util.EnvFile

func bootstrapDependencies() *mux.Router {
	closeApplicationHandler()
	env, err := util.NewEnvFile()
	environment = env
	if err != nil {
		panic(err)
	}
	apiKeyMiddleware := middleware.NewApiKeyMiddleware(environment.GetVariable("API_KEY"))
	siteToScrape := "https://www.180.se/nummer/"
	reviewRepository := persistence.NewSyncMapReviewRepository()
	scrapeSvc := service.NewScrapeService(reviewRepository, siteToScrape)
	reviewController := controller.NewReviewController(scrapeSvc)
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	// Routing
	r := mux.NewRouter()
	r.Use(apiKeyMiddleware.CheckAPIKey)
	r.Use(cors)
	r.HandleFunc("/{phoneNumber}", reviewController.LookUpTelephoneNumber()).Methods(http.MethodGet, http.MethodOptions)

	return r
}

func main() {
	r := bootstrapDependencies()
	log.Fatal(http.ListenAndServe(":8810", r))
}
