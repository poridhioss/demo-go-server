# Go Demo Server

A simple HTTP server demonstrating Go interfaces with swappable database backends.

## Overview

This project shows how Go interfaces allow you to swap implementations (PostgresDB vs MySQLDB) without changing the HTTP handler code. Both databases use in-memory maps for demonstration purposes.

## Prerequisites

- Go 1.21+

## Build and Run

```bash
# Build the binary
go build -o server .

# Run the binary
./server

# Or run directly without building
go run .
```

The server starts on `http://localhost:8080`. You should see:
```
Server on :8080
```

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/users` | GET | Get all users |
| `/user?username=X` | GET | Get user by username |
| `/save?username=X&mobile=Y` | GET | Save a new user |

## Testing the Endpoints

Open another terminal and use curl:

```bash
# Get all users
curl http://localhost:8080/users

# Get single user
curl "http://localhost:8080/user?username=rahim"

# Save new user
curl "http://localhost:8080/save?username=rakib&mobile=01999999999"
```

## Demonstrating Interface Usage

The key concept here is that `main.go` uses a `database` interface:

```go
type database interface {
    Get(username string) (User, error)
    Save(username, mobile string)
    GetAll() []User
}

var db database = NewPostgresDB()  // <-- Change this line to switch databases
```

Both `PostgresDB` and `MySQLDB` implement this interface. To demonstrate, try switching between them:

### Step 1: Run with PostgresDB (default)

```bash
go run .
```

```bash
curl http://localhost:8080/users
# Output: {"users":[{"Username":"rahim","Mobile":"01711111111"},{"Username":"karim","Mobile":"01822222222"}]}
```

Server logs will show: `Running Query in PostgresDB GetAll method`

### Step 2: Switch to MySQLDB

Edit `main.go` line 33 - change:
```go
var db database = NewPostgresDB()
```
to:
```go
var db database = NewMySQLDB()
```

### Step 3: Run again and observe different data

```bash
go run .
```

```bash
curl http://localhost:8080/users
# Output: {"users":[{"Username":"jamal","Mobile":"01933333333"},{"Username":"hasina","Mobile":"01544444444"}]}
```

Server logs will show: `Running MySQLDB Query in GetAll method`

**The HTTP handler code didn't change at all** - only the concrete implementation was swapped. This is the power of Go interfaces!

## Project Structure

```
demo-go-server/
├── main.go        # HTTP server, handlers, and interface definition
├── mysql.go       # MySQLDB implementation
├── postgres.go    # PostgresDB implementation
├── go.mod         # Go module file
├── endpoints.rest # REST Client file for testing endpoints
└── README.md      # This file
```

