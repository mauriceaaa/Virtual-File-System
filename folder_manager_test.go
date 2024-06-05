package main

import (
	"testing"
)

func TestCreateFolder(t *testing.T) {
	userManager := NewUserManager()
	folderManager := NewFolderManager(userManager)
	userManager.AddUser("testuser")

	err := folderManager.CreateFolder("testuser", "testfolder")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = folderManager.CreateFolder("testuser", "testfolder")
	if err == nil {
		t.Fatalf("Expected error for duplicate folder, got none")
	}

	err = folderManager.CreateFolder("nonexistentuser", "testfolder")
	if err == nil {
		t.Fatalf("Expected error for non-existent user, got none")
	}
}

func TestListFolders(t *testing.T) {
	userManager := NewUserManager()
	folderManager := NewFolderManager(userManager)
	userManager.AddUser("testuser")

	folderManager.CreateFolder("testuser", "folder1")
	folderManager.CreateFolder("testuser", "folder2")

	folders, err := folderManager.ListFolders("testuser", "name", "asc")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(folders) != 2 {
		t.Fatalf("Expected 2 folders, got %d", len(folders))
	}
}
