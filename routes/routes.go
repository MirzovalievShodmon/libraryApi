package routes

import (
	"net/http"

	"github.com/MirzovalievShodmon/libraryApi/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("POST /books", handlers.CreateBookHandler)
	http.HandleFunc("GET /books", handlers.GetBooksHandler)
	http.HandleFunc("GET /books/available", handlers.GetAvailableBooksHandler)
	http.HandleFunc("GET /books/search", handlers.SearchBooksHandler)
	http.HandleFunc("GET /books/{id}", handlers.GetBookByIDHandler)
	http.HandleFunc("PUT /books/{id}", handlers.UpdateBookByIDHandler)
	http.HandleFunc("DELETE /books/{id}", handlers.DeleteBookByIDHandler)
	http.HandleFunc("POST /books/{id}/borrow", handlers.BorrowBookByIDHandler)
	http.HandleFunc("POST /books/{id}/return", handlers.ReturnBookByIDHandler)

	http.HandleFunc("POST /users", handlers.CreateUserHandler)
	http.HandleFunc("GET /users", handlers.GetUsersHandler)

}
