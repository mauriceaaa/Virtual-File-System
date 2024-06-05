package main

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	userManager := NewUserManager()
	folderManager := NewFolderManager(userManager)
	fileManager := NewFileManager(userManager)
	userManager.AddUser("testuser")
	folderManager.CreateFolder("testuser", "testfolder")

	err := fileManager.CreateFile("testuser", "testfolder", "testfile", "description")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = fileManager.CreateFile("testuser", "testfolder", "testfile", "description")
	if err == nil {
		t.Fatalf("Expected error for duplicate file, got none")
	}

	err = fileManager.CreateFile("testuser", "nonexistentfolder", "testfile", "description")
	if err == nil {
		t.Fatalf("Expected error for non-existent folder, got none")
	}
}

func TestListFiles(t *testing.T) {
	userManager := NewUserManager()
	folderManager := NewFolderManager(userManager)
	fileManager := NewFileManager(userManager)
	userManager.AddUser("testuser")
	folderManager.CreateFolder("testuser", "testfolder")

	fileManager.CreateFile("testuser", "testfolder", "file1", "description1")
	fileManager.CreateFile("testuser", "testfolder", "file2", "description2")

	files, err := fileManager.ListFiles("testuser", "testfolder", "name", "asc")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("Expected 2 files, got %d", len(files))
	}
}
