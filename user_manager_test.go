package main

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	userManager := NewUserManager()

	err := userManager.AddUser("testuser")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = userManager.AddUser("testuser")
	if err == nil {
		t.Fatalf("Expected error for duplicate user, got none")
	}
}

func TestGetUser(t *testing.T) {
	userManager := NewUserManager()
	userManager.AddUser("testuser")

	user, err := userManager.GetUser("testuser")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if user.Username != "testuser" {
		t.Fatalf("Expected username 'testuser', got %s", user.Username)
	}

	_, err = userManager.GetUser("nonexistentuser")
	if err == nil {
		t.Fatalf("Expected error for non-existent user, got none")
	}
}
