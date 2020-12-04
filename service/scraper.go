package service

import (
	"fmt"
	app_error "phone_lookr/error"
	"phone_lookr/model"
	"phone_lookr/persistence"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type (
	ScrapeService interface {
		Scrape(telephoneNumber string) (*model.Review, error)
	}

	scrapeService struct {
		c          *colly.Collector
		reviewRepo persistence.ReviewRepository
		url        string
	}
)

func NewScrapeService(reviewRepo persistence.ReviewRepository, url string) ScrapeService {
	return &scrapeService{
		c: colly.NewCollector(
			colly.Async(true),
		),
		reviewRepo: reviewRepo,
		url:        url,
	}
}

func (svc *scrapeService) Scrape(telephoneNumber string) (*model.Review, error) {
	review := &model.Review{TelephoneNumber: telephoneNumber}
	r, err := svc.reviewRepo.Find(telephoneNumber)
	if r != nil {
		return r, nil
	}
	if err != nil {
		return nil, err
	}

	svc.scrapeSearchedOnCount(review, err)
	svc.scrapeReportedByCount(review, err)
	svc.scrapeLastSearchedOn(review, err)
	svc.scrapeReviewComments(review, err)

	if err = svc.c.Visit(svc.url + telephoneNumber); err != nil {
		return nil, app_error.SiteCouldNotBeReachedError{
			Url:        svc.url,
			InnerError: err,
		}
	}
	svc.c.Wait()
	if err = svc.reviewRepo.Add(review); err != nil {
		return nil, fmt.Errorf("Could not add review with telNr: %s to storage. %w", review.TelephoneNumber, err)
	}
	return review, nil
}

func (svc *scrapeService) scrapeSearchedOnCount(review *model.Review, callbackError error) {
	svc.c.OnHTML("div.ai-row-1-cell:nth-child(1) > div:nth-child(2)", func(e *colly.HTMLElement) {
		n, err := strconv.Atoi(e.Text)
		if err != nil {
			callbackError = fmt.Errorf("Failed to convert string to integer. Value: %s. %w", e.Text, err)
			return
		}
		review.SearchCount = n
	})
}

func (svc *scrapeService) scrapeReportedByCount(review *model.Review, callbackError error) {
	svc.c.OnHTML("div.ai-row-1-cell:nth-child(2) > div:nth-child(2)", func(e *colly.HTMLElement) {
		n, err := strconv.Atoi(e.Text)
		if err != nil {
			callbackError = fmt.Errorf("Failed to convert string to integer. Value: %s. %w", e.Text, err)
			return
		}
		review.ReportedCount = n
	})
}

func (svc *scrapeService) scrapeLastSearchedOn(review *model.Review, callbackError error) {
	svc.c.OnHTML("div.ai-row-1-cell:nth-child(3) > div:nth-child(2)", func(e *colly.HTMLElement) {
		date := strings.Replace(e.Text, "den ", "", 0)
		callbackError = nil
		review.LastSearch = date
	})
}

func (svc *scrapeService) scrapeReviewComments(review *model.Review, callbackError error) {
	svc.c.OnHTML(".ai-assessment-container", func(e *colly.HTMLElement) {
		comments := make([]model.ReviewComment, 0)
		e.ForEach(".ai-assessment-body", func(i int, h *colly.HTMLElement) {
			comment := model.ReviewComment{
				PostDate: h.ChildText(".ai-assessment-date"),
				Category: model.ParseReviewCommentCategory(h.ChildText(".ai-spam-reason.ai-sr-telephone-seller")),
				Content:  h.ChildText(".ai-comment-text"),
			}
			comments = append(comments, comment)
		})
		callbackError = nil
		review.Comments = comments
	})
}
