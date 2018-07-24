package category

import "net/http"

// Error Declaration
var (
	ErrNameIsRequired = errNameIsRequired{}
	ErrNameIsShorter  = errNameIsShorter{}
)

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsShorter struct{}

func (errNameIsShorter) Error() string {
	return "name must be longer than 5 characters"
}

func (errNameIsShorter) StatusCode() int {
	return http.StatusBadRequest
}
