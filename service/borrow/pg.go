package borrow

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/ntban15/example-go/domain"
)

// pgService implmenter for BorrowRecord serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Validate if user id exist or not
func validateUserIDExist(s *pgService, p *domain.BorrowRecord) error {
	if err := s.db.Where("id = ?", p.UserID).Find(&domain.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrUserIDNotExist
		}
		return err
	}

	return nil
}

// Validate if book id exist or not
func validateBookIDExist(s *pgService, p *domain.BorrowRecord) error {
	if err := s.db.Where("id = ?", p.BookID).Find(&domain.Book{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrBookIDNotExist
		}
		return err
	}

	return nil
}

// Validate if book is available, i.e. not belonging to a record
func validateBookAvailable(s *pgService, p *domain.BorrowRecord) error {
	if err := s.db.Where("book_id = ?", p.BookID).Find(&domain.BorrowRecord{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	return ErrBookNotAvailable
}

// Create implement Create for BorrowRecord service
func (s *pgService) Create(_ context.Context, p *domain.BorrowRecord) error {
	if err := validateUserIDExist(s, p); err != nil {
		return err
	}

	if err := validateBookIDExist(s, p); err != nil {
		return err
	}

	if err := validateBookAvailable(s, p); err != nil {
		return err
	}

	return s.db.Create(p).Error
}

// Update implement Update for BorrowRecord service
func (s *pgService) Update(_ context.Context, p *domain.BorrowRecord) (*domain.BorrowRecord, error) {
	old := domain.BorrowRecord{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if err := validateUserIDExist(s, p); err != nil {
		return nil, err
	}

	if err := validateBookIDExist(s, p); err != nil {
		return nil, err
	}

	if err := validateBookAvailable(s, p); err != nil {
		return nil, err
	}

	old.UserID = p.UserID
	old.BookID = p.BookID
	old.From = p.From
	old.To = p.To

	return &old, s.db.Save(&old).Error
}

// Find implement Find for BorrowRecord service
func (s *pgService) Find(_ context.Context, p *domain.BorrowRecord) (*domain.BorrowRecord, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for BorrowRecord service
func (s *pgService) FindAll(_ context.Context) ([]domain.BorrowRecord, error) {
	res := []domain.BorrowRecord{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for BorrowRecord service
func (s *pgService) Delete(_ context.Context, p *domain.BorrowRecord) error {
	old := domain.BorrowRecord{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}

	return s.db.Delete(old).Error
}
