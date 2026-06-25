package db

import (
	"fmt"
	"log"

	"github.com/MirzovalievShodmon/libraryApi/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var database *sqlx.DB

func ConnectDB() error {
	cfg := configs.LoadConfig()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	database = db

	log.Println("База данных успешно подключена")

	return nil
}

func GetDBConnection() *sqlx.DB {
	return database
}
