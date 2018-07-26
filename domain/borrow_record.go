package domain

import (
	"time"
)

// BorrowRecord describe a record of user borrowing a book
type BorrowRecord struct {
	Model
	UserID UUID      `json:"user_id"`
	BookID UUID      `json:"book_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
