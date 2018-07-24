package domain

// Category describe category of books in system
type Category struct {
	Model
	Name string `json:"name"`
}
