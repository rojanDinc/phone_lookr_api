package controller

import (
	"encoding/json"
	"net/http"
	"phone_lookr/service"

	"github.com/gorilla/mux"
)

type (
	ReviewController interface {
		LookUpTelephoneNumber() func(http.ResponseWriter, *http.Request)
	}

	reviewController struct {
		scrapeService service.ScrapeService
	}
)

func NewReviewController(scrapeSvc service.ScrapeService) ReviewController {
	return &reviewController{
		scrapeService: scrapeSvc,
	}
}

func (ctrl reviewController) LookUpTelephoneNumber() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		review, err := ctrl.scrapeService.Scrape(vars["phoneNumber"])
		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorToMap(err))
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(review)
		}
	}
}
