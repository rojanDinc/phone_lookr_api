package persistence

import (
	"phone_lookr/model"
	"sync"
)

type (
	ReviewRepository interface {
		Add(r *model.Review) error
		Find(telephoneNumber string) (*model.Review, error)
		Remove(telephoneNumber string) error
	}

	reviewRepository struct {
		store *sync.Map
	}
)

func NewSyncMapReviewRepository() ReviewRepository {
	return &reviewRepository{
		store: &sync.Map{},
	}
}

func (repo *reviewRepository) Add(r *model.Review) error {
	repo.store.Store(r.TelephoneNumber, r)
	return nil // Return error if database an error can be created
}

func (repo *reviewRepository) Find(telephoneNumber string) (*model.Review, error) {
	review, ok := repo.store.Load(telephoneNumber)
	if ok == false {
		return nil, nil // Return error if database an error can be created
	}
	return review.(*model.Review), nil
}

func (repo *reviewRepository) Remove(telephoneNumber string) error {
	repo.store.Delete(telephoneNumber)
	return nil // Return error if database an error can be created
}
