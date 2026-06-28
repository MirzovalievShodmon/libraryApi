package repositories

import (
	"database/sql"
	"fmt"

	"github.com/MirzovalievShodmon/libraryApi/db"
	"github.com/MirzovalievShodmon/libraryApi/models"
)

func CreateUser(name, author string) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2)`

	_, err := db.GetDBConnection().Exec(query, name, author)
	if err != nil {
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	users := []models.User{}

	query := `SELECT id, name, email FROM users`

	err := db.GetDBConnection().Select(&users, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения пользователей: %w", err)
	}

	return users, nil
}

func GetUserByID(id int) (models.User, error) {
	var user models.User

	query := `SELECT id, name, email FROM users WHERE id = $1`

	err := db.GetDBConnection().Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("пользователь с id %d не найден", id)
		}

		return models.User{}, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return user, nil
}
