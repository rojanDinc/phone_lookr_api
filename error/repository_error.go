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
	return fmt.Sprintf("Repository error occurred: %v", e.InnerError)
}

func (e RepositoryFindError) Error() string {
	return fmt.Sprintf("Repository error occurred: %v", e.InnerError)
}

func (e RepositoryRemoveError) Error() string {
	return fmt.Sprintf("Repository error occurred: %v", e.InnerError)
}
