package models

import "time"

type BorrowRecord struct {
	ID         int        `json:"id" db:"id"`
	BookID     int        `json:"book_id" db:"book_id"`
	UserID     int        `json:"user_id" db:"user_id"`
	BorrowedAt time.Time  `json:"borrowed_at" db:"borrowed_at"`
	DueDate    time.Time  `json:"due_date" db:"due_date"`
	ReturnedAt *time.Time `json:"returned_at" db:"returned_at"`
}
