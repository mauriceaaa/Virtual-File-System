package main

import (
	"fmt"
)

// User
type User struct {
	Username string             //使用者名稱
	Folders  map[string]*Folder //使用者名稱
}

type UserManager struct {
	users map[string]*User
}

func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[string]*User),
	}
}

// AddUser 新增使用者
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

// GetUser 使用者清單
func (um *UserManager) GetUser(username string) (*User, error) {
	user, exists := um.users[username]
	if !exists {
		return nil, fmt.Errorf("user '%s' does not exist", username)
	}
	return user, nil
}
