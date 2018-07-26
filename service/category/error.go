package category

import "net/http"

// Error Declaration
var (
	ErrNotFound         = errNotFound{}
	ErrNameIsRequired   = errNameIsRequired{}
	ErrNameLength       = errNameLength{}
	ErrNameIsDuplicated = errNameIsDuplicated{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}

func (errNotFound) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNameLength struct{}

func (errNameLength) Error() string {
	return "name must be longer than 5 characters"
}

func (errNameLength) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsDuplicated struct{}

func (errNameIsDuplicated) Error() string {
	return "category name exists"
}

func (errNameIsDuplicated) StatusCode() int {
	return http.StatusBadRequest
}
