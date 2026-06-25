package models

type Book struct {
	ID        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Author    string `json:"author" db:"author"`
	Year      string `json:"year" db:"year"`
	Available bool   `json:"available" db:"available"`
}
