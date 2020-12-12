package persistence

import (
	"phone_lookr/model"
	"testing"
)

func TestAddingAReview(t *testing.T) {
	// Arrange
	repo := NewSyncMapReviewRepository()
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
	// Act
	err := repo.Add(review)
	// Assert
	if err != nil {
		t.Errorf("Could not add review to repository. Reason: %v", err)
	}
}

func TestAddingANilReview(t *testing.T) {
	// Arrange
	repo := NewSyncMapReviewRepository()
	var review *model.Review = nil
	// Act
	err := repo.Add(review)
	// Assert
	if err == nil {
		t.Errorf("Expected an error when adding a nil-value review to repository.")
	}
}

func TestFindingAReview(t *testing.T) {
	// Arrange
	repo := NewSyncMapReviewRepository()
	telNr := "0701234567"
	review := &model.Review{
		TelephoneNumber: telNr,
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
	// Act
	repo.Add(review)
	fetchedReview, err := repo.Find(telNr)
	// Assert
	if err != nil {
		t.Errorf("Got an error finding review with telephone number %s. Reason: %v", telNr, err)
	}
	if fetchedReview == nil {
		t.Errorf("No review was found for telephone number %s", telNr)
	}
}

func TestFindingANonExistentReview(t *testing.T) {
	// Arrange
	repo := NewSyncMapReviewRepository()
	// Act
	review, err := repo.Find("0701234544")
	// Assert
	if err != nil {
		t.Errorf("Failed to find a nil-value review. Reason %v", err)
	}
	if review != nil {
		t.Errorf("Expected not to find a non existent review. Found: %v", review)
	}
}

func TestRemovingReview(t *testing.T) {
	// Arrange
	repo := NewSyncMapReviewRepository()
	telNr := "0701234567"
	review := &model.Review{
		TelephoneNumber: telNr,
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
	// Act
	repo.Add(review)
	err := repo.Remove(telNr)
	// Assert
	if err != nil {
		t.Errorf("Failed to remove review for telephone number %s. Reason: %v", telNr, err)
	}
}

func TestRemovingNonExistentReviewAndFail(t *testing.T) {
	// Arrange
	repo := NewSyncMapReviewRepository()
	// Act
	err := repo.Remove("some number")
	// Assert
	if err == nil {
		t.Error("Expected to get error when removing a non existent review.")
	}
}
