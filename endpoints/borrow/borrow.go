package borrow

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/ntban15/example-go/domain"
	"github.com/ntban15/example-go/service"
)

// CreateData data for CreateBorrowRecord
type CreateData struct {
	UserID domain.UUID `json:"user_id"`
	BookID domain.UUID `json:"book_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// CreateRequest request struct for CreateBorrowRecord
type CreateRequest struct {
	BorrowRecord CreateData `json:"borrow_record"`
}

// CreateResponse response struct for CreateBorrowRecord
type CreateResponse struct {
	BorrowRecord domain.BorrowRecord `json:"borrow_record"`
}

// StatusCode customstatus code for success create BorrowRecord
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a BorrowRecord
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req          = request.(CreateRequest)
			borrowRecord = &domain.BorrowRecord{
				UserID: req.BorrowRecord.UserID,
				BookID: req.BorrowRecord.BookID,
				From:   req.BorrowRecord.From,
				To:     req.BorrowRecord.To,
			}
		)

		err := s.BorrowService.Create(ctx, borrowRecord)
		if err != nil {
			return nil, err
		}

		return CreateResponse{BorrowRecord: *borrowRecord}, nil
	}
}

// FindRequest request struct for Find a BorrowRecord
type FindRequest struct {
	BorrowRecordID domain.UUID
}

// FindResponse response struct for Find a BorrowRecord
type FindResponse struct {
	BorrowRecord *domain.BorrowRecord `json:"borrow_record"`
}

// MakeFindEndPoint make endpoint for find BorrowRecord
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var borrowRecordFind domain.BorrowRecord
		req := request.(FindRequest)
		borrowRecordFind.ID = req.BorrowRecordID

		borrowRecord, err := s.BorrowService.Find(ctx, &borrowRecordFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{BorrowRecord: borrowRecord}, nil
	}
}

// FindAllRequest request struct for FindAll BorrowRecord
type FindAllRequest struct{}

// FindAllResponse request struct for find all BorrowRecord
type FindAllResponse struct {
	BorrowRecords []domain.BorrowRecord `json:"borrow_records"`
}

// MakeFindAllEndpoint make endpoint for find all BorrowRecord
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		borrowRecords, err := s.BorrowService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{BorrowRecords: borrowRecords}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID     domain.UUID `json:"-"`
	UserID domain.UUID `json:"user_id"`
	BookID domain.UUID `json:"book_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	BorrowRecord UpdateData `json:"borrow_record"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	BorrowRecord domain.BorrowRecord `json:"borrow_record"`
}

// MakeUpdateEndpoint make endpoint for update a BorrowRecord
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req          = request.(UpdateRequest)
			borrowRecord = domain.BorrowRecord{
				Model:  domain.Model{ID: req.BorrowRecord.ID},
				UserID: req.BorrowRecord.UserID,
				BookID: req.BorrowRecord.BookID,
				From:   req.BorrowRecord.From,
				To:     req.BorrowRecord.To,
			}
		)

		res, err := s.BorrowService.Update(ctx, &borrowRecord)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{BorrowRecord: *res}, nil
	}
}

// DeleteRequest request struct for delete a BorrowRecord
type DeleteRequest struct {
	BorrowRecordID domain.UUID
}

// DeleteResponse response struct for Find a BorrowRecord
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a BorrowRecord
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			borrowRecordFind = domain.BorrowRecord{}
			req              = request.(DeleteRequest)
		)
		borrowRecordFind.ID = req.BorrowRecordID

		err := s.BorrowService.Delete(ctx, &borrowRecordFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
