package main

import "errors"

// Struct with Map
type PostgresDB struct {
	users map[string]User
}

func NewPostgresDB() PostgresDB {
	return PostgresDB{
		users: map[string]User{
			"rahim": {Username: "rahim", Mobile: "01711111111"},
			"karim": {Username: "karim", Mobile: "01822222222"},
		},
	}
}

// Method
func (p PostgresDB) Get(username string) (User, error) {
	println("Running Query in PostgresDB Get method \n")
	if user, ok := p.users[username]; ok {
		return user, nil
	}
	return User{}, errors.New("user not found in PostgresDB")
}

// Method
func (p PostgresDB) GetAll() []User {
	println("Running Query in PostgresDB GetAll method \n")
	users := []User{}
	for _, user := range p.users {
		users = append(users, user)
	}
	return users
}

// Method
func (p PostgresDB) Save(username, mobile string) {
	p.users[username] = User{Username: username, Mobile: mobile}
}
