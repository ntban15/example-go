package borrow

import "net/http"

// Error declaration
var (
	ErrUserIDNotExist   = errUserIDNotExist{}
	ErrBookIDNotExist   = errBookIDNotExist{}
	ErrBookNotAvailable = errBookNotAvailable{}
)

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
