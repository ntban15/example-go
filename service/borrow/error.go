package borrow

import "net/http"

// Error declaration
var (
	ErrNotFound         = errNotFound{}
	ErrUserIDNotExist   = errUserIDNotExist{}
	ErrBookIDNotExist   = errBookIDNotExist{}
	ErrBookNotAvailable = errBookNotAvailable{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}

func (errNotFound) StatusCode() int {
	return http.StatusBadRequest
}

type errUserIDNotExist struct{}

func (errUserIDNotExist) Error() string {
	return "user id does not exist"
}

func (errUserIDNotExist) StatusCode() int {
	return http.StatusBadRequest
}

type errBookIDNotExist struct{}

func (errBookIDNotExist) Error() string {
	return "book id does not exist"
}

func (errBookIDNotExist) StatusCode() int {
	return http.StatusBadRequest
}

type errBookNotAvailable struct{}

func (errBookNotAvailable) Error() string {
	return "book not available for lending"
}

func (errBookNotAvailable) StatusCode() int {
	return http.StatusBadRequest
}
