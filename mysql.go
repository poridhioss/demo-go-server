package main

import "errors"

// Struct with Map
type MySQLDB struct {
	users map[string]User
}

func NewMySQLDB() MySQLDB {
	return MySQLDB{
		users: map[string]User{
			"jamal":  {Username: "jamal", Mobile: "01933333333"},
			"hasina": {Username: "hasina", Mobile: "01544444444"},
		},
	}
}

// Method
func (m MySQLDB) Get(username string) (User, error) {
	println("Running MySQLDB Query in Get method \n")
	if user, ok := m.users[username]; ok {
		return user, nil
	}
	return User{}, errors.New("user not found in MySQLDB")
}

// Method
func (m MySQLDB) GetAll() []User {
	println("Running MySQLDB Query in GetAll method \n")
	users := []User{}
	for _, user := range m.users {
		users = append(users, user)
	}
	return users
}

// Method
func (m MySQLDB) Save(username, mobile string) {
	m.users[username] = User{Username: username, Mobile: mobile}
}
