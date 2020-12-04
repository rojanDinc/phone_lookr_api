package app_error

import "fmt"

type SiteCouldNotBeReachedError struct {
	Url        string
	InnerError error
}

func (e SiteCouldNotBeReachedError) Error() string {
	return fmt.Sprintf("Site %s could not be reached at this moment. %w", e.Url, e.InnerError)
}
