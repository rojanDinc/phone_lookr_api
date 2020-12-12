package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"phone_lookr/model"
	"phone_lookr/service"
	"testing"
)

type scrapeServiceMock struct {
	review *model.Review
	err    error
}

func (s *scrapeServiceMock) Scrape(telephoneNumber string) (*model.Review, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.review, nil
}

func newServiceMockWithReview(r *model.Review) service.ScrapeService {
	return &scrapeServiceMock{
		review: r,
		err:    nil,
	}
}

func newServiceMockWithError(err error) service.ScrapeService {
	return &scrapeServiceMock{
		review: nil,
		err:    err,
	}
}

func TestLookUpTelephoneNumber(t *testing.T) {
	// Arrange
	telNr := "0701234567"
	review := &model.Review{
		TelephoneNumber: "0701234567",
		LastSearch:      "2020-12-12",
		ReportedCount:   12,
		SearchCount:     123,
		Comments: []model.ReviewComment{
			{
				PostDate: "2020-12-01",
				Content:  "Lorem ipsum dolor sit.",
				Category: model.Sales,
			},
		},
	}
	expectedStatusCode := http.StatusOK
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/"+telNr, nil)
	svc := newServiceMockWithReview(review)
	ctrl := NewReviewController(svc)
	// Act
	ctrl.LookUpTelephoneNumber()(res, req)
	// Assert
	actualStatusCode := res.Result().StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("Got unexpected status code: %d, expected: %d", actualStatusCode, expectedStatusCode)
	}
}

func TestFailLookUpTelephoneNumber(t *testing.T) {
	// Arrange
	telNr := "0701234567"
	err := errors.New("Error mock")
	expectedStatusCode := http.StatusInternalServerError
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/"+telNr, nil)
	svc := newServiceMockWithError(err)
	ctrl := NewReviewController(svc)
	// Act
	fn := ctrl.LookUpTelephoneNumber()
	fn(res, req)
	// Assert
	actualStatusCode := res.Result().StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("Got unexpected status code: %d, expected: %d", actualStatusCode, expectedStatusCode)
	}
}
