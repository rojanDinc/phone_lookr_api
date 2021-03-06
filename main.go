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

func bootstrapDependencies() *mux.Router {
	closeApplicationHandler()
	// Config
	fr := util.NewFileReader(".env")
	err := util.LoadDotEnvFile(fr)
	if err != nil {
		panic(err)
	}
	siteToScrape := "https://www.180.se/nummer/"

	// Repositories
	reviewRepository := persistence.NewSyncMapReviewRepository()

	// Services
	scrapeSvc := service.NewScrapeService(reviewRepository, siteToScrape)

	// Middleware
	apiKeyMiddleware := middleware.NewApiKeyMiddleware(os.Getenv("API_KEY"))

	// Controllers
	reviewController := controller.NewReviewController(scrapeSvc)

	// Routing
	router := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	router.Use(cors)
	router.Use(apiKeyMiddleware.CheckAPIKey)
	router.HandleFunc("/{phoneNumber}", reviewController.LookUpTelephoneNumber()).Methods(http.MethodGet, http.MethodOptions)

	return router
}

func main() {
	r := bootstrapDependencies()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8810"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
