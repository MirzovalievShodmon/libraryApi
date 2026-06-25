package repositories

import (
	"database/sql"
	"fmt"

	"github.com/MirzovalievShodmon/libraryApi/db"
	"github.com/MirzovalievShodmon/libraryApi/models"
)

func CreateBook(title, author string, year int, available bool) error {
	query := `INSERT INTO books (title, author, year, available) VALUES ($1, $2, $3, $4)`

	_, err := db.GetDBConnection().Exec(query, title, author, year, available)
	if err != nil {
		return fmt.Errorf("ошибка создания книги: %w", err)
	}

	return nil
}

func GetAllBooks() ([]models.Book, error) {
	books := []models.Book{}

	query := `SELECT id, title, author, year, available FROM books`

	err := db.GetDBConnection().Select(&books, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения книг: %w", err)
	}

	return books, nil
}

func GetBookByID(id int) (models.Book, error) {
	var book models.Book

	query := `SELECT id, title, author, year, available FROM books WHERE id = $1`

	err := db.GetDBConnection().Get(&book, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, fmt.Errorf("книга с id %d не найдена", id)
		}

		return models.Book{}, fmt.Errorf("ошибка получения книги: %w", err)
	}

	return book, nil
}

func UpdateBookByID(id int, title, author string, year int, available bool) error {
	query := `UPDATE books
	         SET title = $1, author = $2, year = $3, available = $4
			 WHERE id = $5`

	res, err := db.GetDBConnection().Exec(query, title, author, year, available, id)
	if err != nil {
		return fmt.Errorf("ошибка обновления книги: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить результат обновления: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("книга с id %d не найдена", id)
	}

	return nil
}

func DeleteBookByID(id int) error {
	query := `DELETE FROM books WHERE id = $1`

	res, err := db.GetDBConnection().Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления книги: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить результат удаления: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("не удалось получить результат удаления: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("книга с id %d не найдена", id)
	}

	return nil
}

func UpdateBookAvailabilityByID(id int, available bool) error {
	query := `UPDATE books SET available = $1 WHERE id = $2`

	res, err := db.GetDBConnection().Exec(query, available, id)
	if err != nil {
		return fmt.Errorf("ошибка обновления статуса книги: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить результат обновления статуса книги: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("книга с id %d не найдена", id)
	}

	return nil
}

func GetAvailableBooks() ([]models.Book, error) {
	books := []models.Book{}

	query := `SELECT id, title, author, year, available FROM books WHERE available = true`

	err := db.GetDBConnection().Select(&books, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения доступных книг: %w", err)
	}

	return books, nil
}

func SearchBooks(title, author string) ([]models.Book, error) {
	books := []models.Book{}

	query := `
        SELECT id, title, author, year, available
        FROM books
        WHERE ($1 = '' OR title ILIKE '%' || $1 || '%')
          AND ($2 = '' OR author ILIKE '%' || $2 || '%')
`

	err := db.GetDBConnection().Select(&books, query, title, author)
	if err != nil {
		return nil, fmt.Errorf("ошибка поиска книг: %w", err)
	}

	return books, nil
}
