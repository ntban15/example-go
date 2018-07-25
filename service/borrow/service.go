package borrow

import (
	"context"

	"github.com/ntban15/example-go/domain"
)

// Service interface for borrow
type Service interface {
	Create(ctx context.Context, p *domain.BorrowRecord) error
	Update(ctx context.Context, p *domain.BorrowRecord) (*domain.BorrowRecord, error)
	Find(ctx context.Context, p *domain.BorrowRecord) (*domain.BorrowRecord, error)
	FindAll(ctx context.Context) ([]domain.BorrowRecord, error)
	Delete(ctx context.Context, p *domain.BorrowRecord) error
}
