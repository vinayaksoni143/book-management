# Book Management API

A RESTful API service built with Go for managing books. This service provides endpoints to perform CRUD (Create, Read, Update, Delete) operations on books.

## Features

- Get all books
- Get book by ID
- Create new books
- Update existing books
- Delete books

## Technologies Used

- Go
- Gorilla Mux (Router)
- GORM (ORM)
- MySQL Database

## API Endpoints

### Get All Books
`GET /books`
### Get Book by ID
`GET /books/{bookId}`
### Create Book
`POST /books`
### Update Book
`PUT /books/{bookId}`
### Delete Book
`DELETE /books/{bookId}`