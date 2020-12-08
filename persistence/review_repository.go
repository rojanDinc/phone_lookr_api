package persistence

import (
	"errors"
	"fmt"
	app_error "phone_lookr/error"
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
	if r == nil {
		return app_error.RepositoryAddError{
			InnerError: errors.New("The provided review is nil"),
		}
	}
	repo.store.Store(r.TelephoneNumber, r)
	return nil
}

func (repo *reviewRepository) Find(telephoneNumber string) (*model.Review, error) {
	review, ok := repo.store.Load(telephoneNumber)
	if ok == false {
		return nil, nil
	}
	return review.(*model.Review), nil
}

func (repo *reviewRepository) Remove(telephoneNumber string) error {
	review, err := repo.Find(telephoneNumber)
	if err != nil {
		return app_error.RepositoryRemoveError{
			InnerError: err,
		}
	}
	if review == nil {
		return app_error.RepositoryRemoveError{
			InnerError: fmt.Errorf(`A review for the telephone number "%s" does not exist`, telephoneNumber),
		}
	}
	repo.store.Delete(telephoneNumber)
	return nil
}
