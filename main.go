package main

import (
	"log"
	"net/http"

	"github.com/MirzovalievShodmon/libraryApi/configs"
	"github.com/MirzovalievShodmon/libraryApi/db"
	"github.com/MirzovalievShodmon/libraryApi/middlewares"
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

	//Берёт все routes и добавляет к ним логирование.
	handler := middlewares.LoggerMiddleware(http.DefaultServeMux)

	log.Println("Library API запущен на http://localhost:" + cfg.ServerPort)

	err = http.ListenAndServe(":"+cfg.ServerPort, handler)
	if err != nil {
		log.Println("Ошибка запуска сервера:", err)
	}
}
