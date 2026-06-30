package services

import (
	"fmt"

	"github.com/MirzovalievShodmon/libraryApi/models"
	"github.com/MirzovalievShodmon/libraryApi/repositories"
)

func GetAllBorrowRecords() ([]models.BorrowRecord, error) {
	records, err := repositories.GetAllBorrowRecords()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить историю выдачи книг: %w", err)
	}

	return records, nil
}

func GetActiveBorrowRecors() ([]models.BorrowRecord, error) {
	records, err := repositories.GetActiveBorrowRecors()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить активные выдачи книг: %w", err)
	}

	return records, nil
}
