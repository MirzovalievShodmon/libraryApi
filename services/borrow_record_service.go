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

