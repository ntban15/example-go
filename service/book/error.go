package book

import "net/http"

// Error Declaration
var (
	ErrCategoryIDNotExist = errCategoryIDNotExist{}
	ErrNameLength         = errNameLength{}
	ErrDescriptionLength  = errDescriptionLength{}
)

type errCategoryIDNotExist struct{}

func (errCategoryIDNotExist) Error() string {
	return "category id does not exist"
}

func (errCategoryIDNotExist) StatusCode() int {
	return http.StatusBadRequest
}

type errNameLength struct{}

func (errNameLength) Error() string {
	return "book name must be longer than 5 characters"
}

func (errNameLength) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionLength struct{}

func (errDescriptionLength) Error() string {
	return "book description must be longer than 5 characters"
}

func (errDescriptionLength) StatusCode() int {
	return http.StatusBadRequest
}
