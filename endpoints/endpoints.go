package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/ntban15/example-go/service"

	"github.com/ntban15/example-go/endpoints/book"
	"github.com/ntban15/example-go/endpoints/borrow"
	"github.com/ntban15/example-go/endpoints/category"
	"github.com/ntban15/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint

	FindCategory    endpoint.Endpoint
	FindAllCategory endpoint.Endpoint
	CreateCategory  endpoint.Endpoint
	UpdateCategory  endpoint.Endpoint
	DeleteCategory  endpoint.Endpoint

	FindBook    endpoint.Endpoint
	FindAllBook endpoint.Endpoint
	CreateBook  endpoint.Endpoint
	UpdateBook  endpoint.Endpoint
	DeleteBook  endpoint.Endpoint

	FindBorrowRecord    endpoint.Endpoint
	FindAllBorrowRecord endpoint.Endpoint
	CreateBorrowRecord  endpoint.Endpoint
	UpdateBorrowRecord  endpoint.Endpoint
	DeleteBorrowRecord  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),

		FindCategory:    category.MakeFindEndPoint(s),
		FindAllCategory: category.MakeFindAllEndpoint(s),
		CreateCategory:  category.MakeCreateEndpoint(s),
		UpdateCategory:  category.MakeUpdateEndpoint(s),
		DeleteCategory:  category.MakeDeleteEndpoint(s),

		FindBook:    book.MakeFindEndPoint(s),
		FindAllBook: book.MakeFindAllEndpoint(s),
		CreateBook:  book.MakeCreateEndpoint(s),
		UpdateBook:  book.MakeUpdateEndpoint(s),
		DeleteBook:  book.MakeDeleteEndpoint(s),

		FindBorrowRecord:    borrow.MakeFindEndPoint(s),
		FindAllBorrowRecord: borrow.MakeFindAllEndpoint(s),
		CreateBorrowRecord:  borrow.MakeCreateEndpoint(s),
		UpdateBorrowRecord:  borrow.MakeUpdateEndpoint(s),
		DeleteBorrowRecord:  borrow.MakeDeleteEndpoint(s),
	}
}
