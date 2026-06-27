package services

import (
	"fmt"
	"strings"

	"github.com/MirzovalievShodmon/libraryApi/models"
	"github.com/MirzovalievShodmon/libraryApi/repositories"
)

func CreateBook(title, author string, year int, available bool) error {
	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)

	if title == "" {
		return fmt.Errorf("название книги не может быть пустым")
	}

	if author == "" {
		return fmt.Errorf("автор книги не может быть пустым")
	}

	if year < 0 {
		return fmt.Errorf("год книги не может быть отрицательным")
	}

	err := repositories.CreateBook(title, author, year, available)
	if err != nil {
		return fmt.Errorf("не удалось создать книгу: %w", err)
	}

	return nil
}

func GetAllBooks() ([]models.Book, error) {
	books, err := repositories.GetAllBooks()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить книги: %w", err)
	}

	return books, nil
}

func GetBookByID(id int) (models.Book, error) {
	if id <= 0 {
		return models.Book{}, fmt.Errorf("id книги должен быть положительным числом")
	}

	book, err := repositories.GetBookByID(id)
	if err != nil {
		return models.Book{}, fmt.Errorf("не удалось получить книгу по id: %w", err)
	}

	return book, nil
}

func UpdateBookByID(id int, title, author string, year int, available bool) error {
	if id <= 0 {
		return fmt.Errorf("id книги должен быть положительным числом")
	}

	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)

	if title == "" {
		return fmt.Errorf("название книги не может быть пустым")
	}

	if author == "" {
		return fmt.Errorf("автор книги не может быть пустым")
	}

	if year < 0 {
		return fmt.Errorf("год книги не может быть отрицательным")
	}

	err := repositories.UpdateBookByID(id, title, author, year, available)
	if err != nil {
		return fmt.Errorf("не удалось обновить книгу: %w", err)
	}

	return nil
}

func DeleteBookByID(id int) error {
	if id <= 0 {
		return fmt.Errorf("id книги должен быть положительным числом")
	}

	err := repositories.DeleteBookByID(id)
	if err != nil {
		return fmt.Errorf("не удалось удалить книгу: %w", err)
	}

	return nil
}

func BorrowBookByID(bookID, userID int) error {
	if bookID <= 0 {
		return fmt.Errorf("id книги должен быть положительным числом")
	}

	if userID <= 0 {
		return fmt.Errorf("id пользователя должен быть положительным числом")
	}

	book, err := repositories.GetBookByID(bookID)
	if err != nil {
		return fmt.Errorf("не удалось получить книгу: %w", err)
	}

	if !book.Available {
		return fmt.Errorf("книга уже выдана")
	}

	err = repositories.CreateBorrowRecord(bookID, userID)
	if err != nil {
		return fmt.Errorf("не удалось создать запись выдачи книги: %w", err)
	}

	err = repositories.UpdateBookAvailabilityByID(bookID, false)
	if err != nil {
		return fmt.Errorf("не удалось выдать книгу: %w", err)
	}

	return nil
}

func ReturnBookByID(bookID int) error {
	if bookID <= 0 {
		return fmt.Errorf("id книги должен быть положительным числом")
	}

	book, err := repositories.GetBookByID(bookID)
	if err != nil {
		return fmt.Errorf("не удалось получить книгу: %w", err)
	}

	if book.Available {
		return fmt.Errorf("книга уже находится в библиотеке")
	}

	err = repositories.CloseBorrowRecord(bookID)
	if err != nil {
		return fmt.Errorf("не удалось закрыть запись на выдачи книги: %w", err)
	}

	err = repositories.UpdateBookAvailabilityByID(bookID, true)
	if err != nil {
		return fmt.Errorf("не удалось вернуть книгу: %w", err)
	}

	return nil
}

func GetAvailableBooks() ([]models.Book, error) {
	books, err := repositories.GetAvailableBooks()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить доступные книги: %w", err)
	}

	return books, nil
}

func SearchBooks(title, author string) ([]models.Book, error) {
	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)

	if title == "" && author == "" {
		return nil, fmt.Errorf("нужно указать title или author для поиска")
	}

	books, err := repositories.SearchBooks(title, author)
	if err != nil {
		return nil, fmt.Errorf("не удалось найти книги: %w", err)
	}

	return books, nil
}
