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

func CreateBorrowRecord(bookID, userID int) error {
	query := `
	    INSERT INTO borrow_records (book_id, user_id, due_date)
	    VALUES ($1, $2, NOW() + '14 days')
	`

	_, err := db.GetDBConnection().Exec(query, bookID, userID)
	if err != nil {
		return fmt.Errorf("ошибка создания записи выдачи книги: %w", err)
	}

	return nil
}

func CloseBorrowRecord(bookID int) error {
	query := `
        UPDATE borrow_records
	    SET returned_at = NOW()
		WHERE id = (
		    SELECT id
			FROM borrow_records
			WHERE book_id = $1 AND returned_at IS NULL
			ORDER BY borrowed_at DESC
			LIMIT 1
		)
`

	res, err := db.GetDBConnection().Exec(query, bookID)
	if err != nil {
		return fmt.Errorf("ошибка закрытия записи выдачи книги: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить результат закрытия выдачи книги: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("активная запись выдачи для книги с id %d не найдена", bookID)
	}

	return nil
}

func GetActiveBorrowRecords() ([]models.BorrowRecord, error) {
	records := []models.BorrowRecord{}
	query := `
        SELECT id, book_id, user_id, borrowed_at, due_date, returned_at
        FROM borrow_records
        WHERE returned_at IS NULL
        ORDER BY borrowed_at DESC
`

	err := db.GetDBConnection().Select(&records, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения активных выдач книг: %w", err)
	}

	return records, nil
}

func BorrowBookWithTransaction(bookID, userID int) error {
	tx, err := db.GetDBConnection().Beginx()
	if err != nil {
		return fmt.Errorf("ошибка начала транзакции: %w", err)
	}

	defer tx.Rollback()

	queryCreateRecord := `
		INSERT INTO borrow_records ( book_id, user_id, due_date )
		VALUES ($1, $2, NOW() + INTERVAL '14 days')
		`

	_, err = tx.Exec(queryCreateRecord, bookID, userID)
	if err != nil {
		return fmt.Errorf("ошибка создания записи выдачи книги: %w", err)
	}

	queryUpdateBook := `
	UPDATE books
	SET available = false
	WHERE id = $1
	`

	res, err := tx.Exec(queryUpdateBook, bookID)
	if err != nil {
		return fmt.Errorf("ошибка обновления статуса книги: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить результат обновления книги: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("книга с id %d не найдена", bookID)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("ошибка подтверждения транзакции: %w", err)
	}

	return nil
}

func ReturnBookWithTransaction(bookID int) error {
	tx, err := db.GetDBConnection().Beginx()
	if err != nil {
		return fmt.Errorf(" ошибка начало транзакции: %w", err)
	}

	defer tx.Rollback()

	queryClosed := `
	    UPDATE borrow_records
	    SET returned_at = NOW()
	    WHERE id = (
	        SELECT id
	        FROM borrow_records
	        WHERE book_id = $1 AND returned_at IS NULL
	        ORDER BY borrowed_at DESC
	        LIMIT 1
	    )
	`

	res, err := tx.Exec(queryClosed, bookID)
	if err != nil {
		return fmt.Errorf("ошибка закрытия записи выдачи книги: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить результат закрытия записи выдачи книги: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("активная запись выдачи для книги с id %d не найдена", bookID)
	}

	queryUpdateBook := `
	UPDATE books
	SET available = true
	WHERE id = $1
	`

	res, err = tx.Exec(queryUpdateBook, bookID)
	if err != nil {
		return fmt.Errorf("ошибка обновления статуса книги: %w", err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить результат обновления книги: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("книга с id %d не найдена", bookID)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("ошибка подтверждения транзакции: %w", err)
	}

	return nil
}
