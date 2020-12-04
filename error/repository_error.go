package app_error

import "fmt"

type (
	RepositoryAddError struct {
		InnerError error
	}
	RepositoryFindError struct {
		InnerError error
	}
	RepositoryRemoveError struct {
		InnerError error
	}
)

func (e RepositoryAddError) Error() string {
	return fmt.Sprintf("Persistence error occured: %w", e.InnerError)
}
