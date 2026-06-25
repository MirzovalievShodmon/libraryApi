package main

import (
	"log"
	"net/http"

	"github.com/MirzovalievShodmon/libraryApi/configs"
	"github.com/MirzovalievShodmon/libraryApi/db"
	"github.com/MirzovalievShodmon/libraryApi/routes"
)

func main() {
	cfg := configs.LoadConfig()

	err := db.ConnectDB()
	if err != nil {
		log.Println(err)
		return
	}

	routes.RegisterRoutes()

	log.Println("Library API запущен на http://localhost:" + cfg.ServerPort)

	err = http.ListenAndServe(":"+cfg.ServerPort, nil)
	if err != nil {
		log.Println("Ошибка запуска сервера:", err)
	}
}
