# Library API

Library API — это REST API проект на Go для управления книгами в библиотеке.

Проект позволяет:

* создавать книги
* получать список всех книг
* получать только доступные книги
* искать книги по названию или автору
* получать книгу по ID
* обновлять книгу
* удалять книгу
* выдавать книгу
* возвращать книгу

## Технологии

В проекте используются:

* Go
* PostgreSQL
* net/http
* sqlx
* lib/pq
* godotenv
* JSON
* REST API
* Postman для тестирования

## Структура проекта

```text
libraryApi/
├── main.go
├── .env
├── .gitignore
├── README.md
├── configs/
│   └── config.go
├── db/
│   └── db.go
├── models/
│   └── book.go
├── repositories/
│   └── book_repository.go
├── services/
│   └── book_service.go
├── handlers/
│   └── book_handler.go
└── routes/
    └── routes.go
```

## Архитектура

Проект разделён на несколько слоёв:

```text
main.go → routes → handlers → services → repositories → PostgreSQL
```

### main.go

`main.go` является точкой входа в приложение.

Он выполняет основные действия:

* загружает конфигурацию
* подключает базу данных
* регистрирует маршруты
* запускает HTTP-сервер

### Configs

Пакет `configs` отвечает за загрузку настроек из файла `.env`.

Например:

* порт сервера
* host базы данных
* port базы данных
* имя пользователя PostgreSQL
* пароль PostgreSQL
* название базы данных

### Routes

Пакет `routes` отвечает за регистрацию HTTP routes.

Например:

```go
http.HandleFunc("GET /books", handlers.GetBooksHandler)
http.HandleFunc("POST /books", handlers.CreateBookHandler)
```

### Handlers

Handlers принимают HTTP-запросы, читают JSON, получают параметры из URL и возвращают JSON-ответ клиенту.

### Services

Services содержат бизнес-логику проекта.

Например:

* название книги не должно быть пустым
* автор книги не должен быть пустым
* год книги не должен быть отрицательным
* нельзя выдать книгу, если она уже выдана
* нельзя вернуть книгу, если она уже находится в библиотеке

### Repositories

Repositories работают с базой данных PostgreSQL и выполняют SQL-запросы:

* INSERT
* SELECT
* UPDATE
* DELETE

## Конфигурация

В проекте используется файл `.env`.

Пример `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=library_db
SERVER_PORT=8080
```

Файл `.env` добавлен в `.gitignore`, потому что в нём могут храниться приватные данные, например пароль от базы данных.

Пример `.gitignore`:

```gitignore
.env
.DS_Store
```

## База данных

Используется база данных PostgreSQL.

Название базы данных:

```text
library_db
```

Таблица книг:

```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    year INT,
    available BOOLEAN DEFAULT TRUE
);
```

## Запуск проекта

1. Клонировать или открыть проект.

2. Установить зависимости:

```bash
go mod tidy
```

3. Создать базу данных PostgreSQL:

```sql
CREATE DATABASE library_db;
```

4. Создать таблицу `books`:

```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    year INT,
    available BOOLEAN DEFAULT TRUE
);
```

5. Создать файл `.env` и добавить настройки:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=library_db
SERVER_PORT=8080
```

6. Запустить проект:

```bash
go run main.go
```

После запуска сервер будет доступен по адресу:

```text
http://localhost:8080
```

## Endpoints

### Проверка сервера

```http
GET /hello
```

Пример ответа:

```json
{
  "message": "Hello from Library API"
}
```

---

### Создать книгу

```http
POST /books
```

Пример JSON:

```json
{
  "title": "Clean Code",
  "author": "Robert Martin",
  "year": 2008,
  "available": true
}
```

Пример ответа:

```json
{
  "message": "Книга успешно создана"
}
```

---

### Получить все книги

```http
GET /books
```

Пример ответа:

```json
[
  {
    "id": 1,
    "title": "Clean Code",
    "author": "Robert Martin",
    "year": 2008,
    "available": true
  }
]
```

Если книг нет, вернётся пустой список:

```json
[]
```

---

### Получить только доступные книги

```http
GET /books/available
```

Этот endpoint возвращает только книги, у которых:

```text
available = true
```

Пример ответа:

```json
[
  {
    "id": 1,
    "title": "Clean Code",
    "author": "Robert Martin",
    "year": 2008,
    "available": true
  }
]
```

Если доступных книг нет, вернётся пустой список:

```json
[]
```

---

### Поиск книг

```http
GET /books/search
```

Поиск можно выполнять по названию книги:

```http
GET /books/search?title=code
```

По автору:

```http
GET /books/search?author=martin
```

По названию и автору вместе:

```http
GET /books/search?title=go&author=donovan
```

Пример ответа:

```json
[
  {
    "id": 1,
    "title": "Clean Code",
    "author": "Robert Martin",
    "year": 2008,
    "available": false
  }
]
```

Если ничего не найдено, вернётся пустой список:

```json
[]
```

Если не указать `title` или `author`, вернётся ошибка:

```json
{
  "error": "нужно указать title или author для поиска"
}
```

---

### Получить книгу по ID

```http
GET /books/{id}
```

Пример:

```http
GET /books/1
```

Пример ответа:

```json
{
  "id": 1,
  "title": "Clean Code",
  "author": "Robert Martin",
  "year": 2008,
  "available": true
}
```

Если ID неправильный:

```json
{
  "error": "неправильный id книги"
}
```

Если ID отрицательный или равен нулю:

```json
{
  "error": "id книги должен быть положительным числом"
}
```

---

### Обновить книгу по ID

```http
PUT /books/{id}
```

Пример:

```http
PUT /books/1
```

Пример JSON:

```json
{
  "title": "Clean Code Updated",
  "author": "Robert Martin",
  "year": 2008,
  "available": true
}
```

Пример ответа:

```json
{
  "message": "Книга успешно обновлена"
}
```

---

### Удалить книгу по ID

```http
DELETE /books/{id}
```

Пример:

```http
DELETE /books/1
```

Пример ответа:

```json
{
  "message": "Книга успешно удалена"
}
```

---

### Выдать книгу

```http
POST /books/{id}/borrow
```

Пример:

```http
POST /books/1/borrow
```

Если книга доступна, она становится недоступной:

```text
available: true → false
```

Пример ответа:

```json
{
  "message": "Книга успешно выдана"
}
```

Если книга уже выдана:

```json
{
  "error": "книга уже выдана"
}
```

---

### Вернуть книгу

```http
POST /books/{id}/return
```

Пример:

```http
POST /books/1/return
```

Если книга была выдана, она снова становится доступной:

```text
available: false → true
```

Пример ответа:

```json
{
  "message": "Книга успешно возвращена"
}
```

Если книга уже находится в библиотеке:

```json
{
  "error": "книга уже находится в библиотеке"
}
```

## Логирование

В проект добавлены простые логи через пакет `log`.

Пример логов:

```text
2026/06/24 03:23:10 База данных успешно подключена
2026/06/24 03:23:10 Library API запущен на http://localhost:8080
2026/06/24 03:23:53 Книга успешно создана: Test Book
2026/06/24 03:26:39 Книга успешно выдана, id: 3
2026/06/24 03:26:47 Книга успешно возвращена, id: 3
```

## Примеры тестирования в Postman

Создать книгу:

```http
POST http://localhost:8080/books
```

Получить все книги:

```http
GET http://localhost:8080/books
```

Получить доступные книги:

```http
GET http://localhost:8080/books/available
```

Поиск по автору:

```http
GET http://localhost:8080/books/search?author=martin
```

Поиск по названию:

```http
GET http://localhost:8080/books/search?title=code
```

Выдать книгу:

```http
POST http://localhost:8080/books/1/borrow
```

Вернуть книгу:

```http
POST http://localhost:8080/books/1/return
```

## Возможные улучшения

В будущем можно добавить:

* пользователей
* историю выдачи книг
* срок возврата книги
* штраф за просрочку или плохое состояние книги
* JWT авторизацию
* роли администратора и пользователя
* Swagger документацию
* unit tests
* транзакции
* работу с файлами, например обложки книг
