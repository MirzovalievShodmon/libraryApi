package services

import (
	"fmt"
	"strings"

	"github.com/MirzovalievShodmon/libraryApi/models"
	"github.com/MirzovalievShodmon/libraryApi/repositories"
)

func CreateUser(name, email string) error {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return fmt.Errorf("имя пользователя не может быть пустым")
	}

	if email == "" {
		return fmt.Errorf("email пользователя не может быть пустым")
	}

	err := repositories.CreateUser(name, email)
	if err != nil {
		return fmt.Errorf("не удалось создать пользователя: %w", err)
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить пользователей: %w", err)
	}

	return users, err
}

func GetUserByID(id int) (models.User, error) {
	if id <= 0 {
		return models.User{}, fmt.Errorf("id пользовавтеля должен быть положительным числом")
	}

	user, err := repositories.GetUserByID(id)
	if err != nil {
		return models.User{}, fmt.Errorf("не удалось получить пользователя: %w", err)
	}

	return user, nil
}
