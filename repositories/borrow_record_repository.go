package repositories

import (
	"fmt"

	"github.com/MirzovalievShodmon/libraryApi/db"
	"github.com/MirzovalievShodmon/libraryApi/models"
)

func GetAllBorrowRecords() ([]models.BorrowRecord, error) {
	records := []models.BorrowRecord{}

	query := `
	SELECT id, book_id, user_id, borrowed_at, due_date, returned_at
	FROM borrow_records
	ORDER BY id
	`

	err := db.GetDBConnection().Select(&records, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения истории выдачи книг: %w", err)
	}

	return records, err
}
