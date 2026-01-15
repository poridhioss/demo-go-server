package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct
type User struct {
	Username string
	Mobile   string
}

// Request Body
type RequestBody struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
}

// Response Body
type ResponseBody struct {
	Users []User `json:"users"`
}

// Interface - both PostgresDB and MySQLDB implement this
type database interface {
	Get(username string) (User, error)
	Save(username, mobile string)
	GetAll() []User
}

var db database = NewMySQLDB() // swap with NewMySQLDB() to switch

func main() {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		user, err := db.Get(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		responseBody := ResponseBody{
			Users: []User{user},
		}

		json.NewEncoder(w).Encode(responseBody)
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := db.GetAll()
		responseBody := ResponseBody{
			Users: users,
		}
		json.NewEncoder(w).Encode(responseBody)
	})

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		mobile := r.URL.Query().Get("mobile")
		db.Save(username, mobile)
		fmt.Fprintf(w, "Saved: %s (%s)", username, mobile)
	})

	fmt.Println("Server on :8080")
	http.ListenAndServe(":8080", nil)
}
