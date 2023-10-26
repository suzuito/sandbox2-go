package usecase

import "fmt"

type RepositoryErrorCode string

var (
	RepositoryErrorCodeNotFound RepositoryErrorCode = "not found"
	RepositoryErrorCodeNotImpl  RepositoryErrorCode = "not impl"
)

type RepositoryError struct {
	EntityURL string
	Message   string
	Code      RepositoryErrorCode
}

func (t *RepositoryError) Error() string {
	if t.Code == RepositoryErrorCodeNotFound {
		return fmt.Sprintf("Not found %s", t.EntityURL)
	}
	return fmt.Sprintf(
		"Error on '%s' %s",
		t.EntityURL,
		t.Message,
	)
}

type SortOrder string

var (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)
