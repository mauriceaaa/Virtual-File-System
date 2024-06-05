package main

import (
	"fmt"
)

type User struct {
	Username string
	Folders  map[string]*Folder
}

type UserManager struct {
	users map[string]*User
}

func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[string]*User),
	}
}

func (um *UserManager) AddUser(username string) error {
	if _, exists := um.users[username]; exists {
		return fmt.Errorf("user '%s' already exists", username)
	}
	um.users[username] = &User{
		Username: username,
		Folders:  make(map[string]*Folder),
	}
	return nil
}

func (um *UserManager) GetUser(username string) (*User, error) {
	user, exists := um.users[username]
	if !exists {
		return nil, fmt.Errorf("user '%s' does not exist", username)
	}
	return user, nil
}
