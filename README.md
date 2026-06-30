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
* создавать пользователей
* получать список пользователей
* выдавать книгу конкретному пользователю
* возвращать книгу
* хранить историю выдачи книг
* получать всю историю выдачи книг
* получать только активные выдачи книг

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
│   ├── book.go
│   ├── user.go
│   └── borrow_record.go
├── repositories/
│   ├── book_repository.go
│   ├── user_repository.go
│   └── borrow_record_repository.go
├── services/
│   ├── book_service.go
│   ├── user_service.go
│   └── borrow_record_service.go
├── handlers/
│   ├── book_handler.go
│   ├── user_handler.go
│   └── borrow_record_handler.go
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
http.HandleFunc("GET /borrow-records", handlers.GetBorrowRecordsHandler)
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
* нельзя выдать книгу несуществующему пользователю

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

### Таблица книг

```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    year INT,
    available BOOLEAN DEFAULT TRUE
);
```

Поля таблицы `books`:

* `id` — уникальный ID книги
* `title` — название книги
* `author` — автор книги
* `year` — год выпуска книги
* `available` — доступна книга или нет

Если:

```text
available = true
```

значит книга находится в библиотеке.

Если:

```text
available = false
```

значит книга выдана пользователю.

### Таблица пользователей

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);
```

Поля таблицы `users`:

* `id` — уникальный ID пользователя
* `name` — имя пользователя
* `email` — email пользователя

### Таблица истории выдачи книг

```sql
CREATE TABLE borrow_records (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL REFERENCES books(id),
    user_id INT NOT NULL REFERENCES users(id),
    borrowed_at TIMESTAMP DEFAULT NOW(),
    due_date TIMESTAMP NOT NULL,
    returned_at TIMESTAMP
);
```

Таблица `borrow_records` хранит историю выдачи книг.

Поля таблицы `borrow_records`:

* `id` — уникальный ID записи выдачи
* `book_id` — какую книгу взяли
* `user_id` — какой пользователь взял книгу
* `borrowed_at` — когда книгу взяли
* `due_date` — до какой даты нужно вернуть книгу
* `returned_at` — когда книгу реально вернули

Если:

```text
returned_at = NULL
```

значит книга ещё не возвращена.

Если:

```text
returned_at != NULL
```

значит книгу уже вернули.

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

5. Создать таблицу `users`:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);
```

6. Создать таблицу `borrow_records`:

```sql
CREATE TABLE borrow_records (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL REFERENCES books(id),
    user_id INT NOT NULL REFERENCES users(id),
    borrowed_at TIMESTAMP DEFAULT NOW(),
    due_date TIMESTAMP NOT NULL,
    returned_at TIMESTAMP
);
```

7. Создать файл `.env` и добавить настройки:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=library_db
SERVER_PORT=8080
```

8. Запустить проект:

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

### Создать пользователя

```http
POST /users
```

Пример JSON:

```json
{
  "name": "Ahmad",
  "email": "ahmad@mail.com"
}
```

Пример ответа:

```json
{
  "message": "Пользователь успешно создан"
}
```

Если имя пустое:

```json
{
  "error": "имя пользователя не может быть пустым"
}
```

Если email пустой:

```json
{
  "error": "email пользователя не может быть пустым"
}
```

---

### Получить всех пользователей

```http
GET /users
```

Пример ответа:

```json
[
  {
    "id": 1,
    "name": "Ahmad",
    "email": "ahmad@mail.com"
  }
]
```

Если пользователей нет, вернётся пустой список:

```json
[]
```

---

### Выдать книгу пользователю

```http
POST /books/{id}/borrow
```

Пример:

```http
POST /books/1/borrow
```

Пример JSON:

```json
{
  "user_id": 1
}
```

При выдаче книги происходит два действия:

```text
books.available: true → false
```

И создаётся новая запись в таблице `borrow_records`.

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

Если пользователя нет:

```json
{
  "error": "не удалось получить пользователя: пользователь с id 999 не найден"
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

При возврате книги происходит два действия:

```text
books.available: false → true
```

И в таблице `borrow_records` поле `returned_at` заполняется текущим временем.

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

---

### Получить всю историю выдачи книг

```http
GET /borrow-records
```

Этот endpoint возвращает все записи выдачи книг:

* книги, которые уже вернули
* книги, которые ещё не вернули

Пример ответа:

```json
[
  {
    "id": 1,
    "book_id": 1,
    "user_id": 1,
    "borrowed_at": "2026-06-27T15:58:33.031002Z",
    "due_date": "2026-07-11T15:58:33.031002Z",
    "returned_at": null
  }
]
```

Если `returned_at = null`, значит книга ещё не возвращена.

Если история выдачи пустая, вернётся:

```json
[]
```

---

### Получить активные выдачи книг

```http
GET /borrow-records/active
```

Этот endpoint возвращает только книги, которые сейчас выданы и ещё не возвращены.

То есть записи, где:

```sql
returned_at IS NULL
```

Пример ответа:

```json
[
  {
    "id": 1,
    "book_id": 1,
    "user_id": 1,
    "borrowed_at": "2026-06-27T15:58:33.031002Z",
    "due_date": "2026-07-11T15:58:33.031002Z",
    "returned_at": null
  }
]
```

Если активных выдач нет, вернётся:

```json
[]
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
2026/06/24 03:30:12 Пользователь успешно создан: ahmad@mail.com
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

Создать пользователя:

```http
POST http://localhost:8080/users
```

Получить пользователей:

```http
GET http://localhost:8080/users
```

Выдать книгу пользователю:

```http
POST http://localhost:8080/books/1/borrow
```

Body:

```json
{
  "user_id": 1
}
```

Вернуть книгу:

```http
POST http://localhost:8080/books/1/return
```

Получить всю историю выдачи:

```http
GET http://localhost:8080/borrow-records
```

Получить активные выдачи:

```http
GET http://localhost:8080/borrow-records/active
```

## Возможные улучшения

В будущем можно добавить:

* проверку email на уникальность
* красивую обработку duplicate email
* штраф за просрочку или плохое состояние книги
* JWT авторизацию
* роли администратора и пользователя
* Swagger документацию
* unit tests
* транзакции для borrow/return
* работу с файлами, например обложки книг
