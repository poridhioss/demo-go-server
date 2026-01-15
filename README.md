# Demo Server

A simple HTTP server demonstrating Go interfaces with swappable database backends.

## Overview

This project shows how Go interfaces allow you to swap implementations (MySQL vs Postgres) without changing the HTTP handler code. Both databases use in-memory maps for demonstration purposes.

## Prerequisites

- Go 1.21+

## How to Run

```bash
# Navigate to the project directory
cd demo-server

# Run the server
go run .
```

The server starts on `http://localhost:8080`

You should see:
```
Server on :8080
```

## How to Test

Open another terminal and use curl to test the endpoints:

### Get All Users

```bash
curl http://localhost:8080/users
```

**Expected Output (MySQL - default):**
```json
{"users":[{"Username":"jamal","Mobile":"01933333333"},{"Username":"hasina","Mobile":"01544444444"}]}
```

### Get Single User

```bash
curl "http://localhost:8080/user?username=jamal"
```

**Expected Output:**
```json
{"users":[{"Username":"jamal","Mobile":"01933333333"}]}
```

### Get Non-Existent User

```bash
curl "http://localhost:8080/user?username=unknown"
```

**Expected Output:**
```
user not found in MySQLDB
```

### Save New User

```bash
curl "http://localhost:8080/save?username=rakib&mobile=01999999999"
```

**Expected Output:**
```
Saved: rakib (01999999999)
```

Then verify:
```bash
curl "http://localhost:8080/user?username=rakib"
```

## Switching Database

To switch from MySQL to Postgres, edit `main.go` line 33:

```go
// Change from:
var db database = NewMySQLDB()

// To:
var db database = NewPostgresDB()
```

Postgres has different seed data (rahim, karim instead of jamal, hasina).

## Project Structure

```
demo-server/
├── main.go       # HTTP server and handlers
├── mysql.go      # MySQL database implementation
├── postgres.go   # Postgres database implementation
├── go.mod        # Go module file
└── endpoints.md  # API endpoint examples
```

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/users` | GET | Get all users |
| `/user?username=X` | GET | Get user by username |
| `/save?username=X&mobile=Y` | GET | Save a new user |

