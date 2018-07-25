package service

import (
	"github.com/ntban15/example-go/service/book"
	"github.com/ntban15/example-go/service/borrow"
	"github.com/ntban15/example-go/service/category"
	"github.com/ntban15/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	BorrowService   borrow.Service
}
