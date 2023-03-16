# Gorilla/Mux REST API

This is a sample RESTful API built using Gorilla/Mux in Go programming language. This API allows you to perform CRUD operations on a user database.

## Prerequisites

Go 1.14 or later
Gorilla/Mux package
PostgreSQL database

## Installation

Clone this repository
Install dependencies

```go
go mod download
```

Create a .env file at the root of the project and set the following environment variables:

```makefile
DB_HOST=<database-host>
DB_PORT=<database-port>
DB_USER=<database-username>
DB_PASSWORD=<database-password>
DB_NAME=<database-name>
```

Start the server

```go
go run cmd/server/main.go
```

The server will be running at <http://localhost:8080>.

## Endpoints

| Method| Path  |Description|
| ----- | ----  |-----------|
| GET   | /users|Get all users|
| GET   | /users/{id}|  Get a single user by ID|
| POST  | /users |Create a new user|
| PUT   | /users/{id} | Update an existing user by ID|
| DELETE| /users/{id}| Delete a user by ID|

### Request/Response Format

All requests and responses are in JSON format.

Get all users

Request

```perl
GET /users
```

Response

```perl

Status: 200 OK
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "johndoe@example.com",
    "age": 30
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "janesmith@example.com",
    "age": 25
  }
]
```

Get a single user by ID

#### Request

```bash
GET /users/1
```

### Response

```perl
Status: 200 OK
{
  "id": 1,
  "name": "John Doe",
  "email": "johndoe@example.com",
  "age": 30
}
```

Create a new user

### Request

```bash
POST /users
{
  "name": "Alice Brown",
  "email": "alicebrown@example.com",
  "age": 35
}

```

### Response

```css
Status: 201 Created
{
  "id": 3,
  "name": "Alice Brown",
  "email": "alicebrown@example.com",
  "age": 35
}
```

Update an existing user by ID

### Request

```bash
PUT /users/3

{
  "name": "Alice Green",
  "email": "alicegreen@example.com",
  "age": 40
}
```

### Response

```css
Status: 200 OK
{
  "id": 3,
  "name": "Alice Green",
  "email": "alicegreen@example.com",
  "age": 40
}
```

Delete a user by ID

### Request

```bash
DELETE /users/3
```

### Response

``` yaml
Status: 204 No Content
```

## Helpful Links

[Getting Started Tutorial](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
