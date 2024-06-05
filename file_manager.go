package main

import (
	"fmt"
	"sort"
	"time"
)

type File struct {
	Name        string
	Description string
	CreatedAt   time.Time
}

type FileManager struct {
	userManager *UserManager
}

func NewFileManager(userManager *UserManager) *FileManager {
	return &FileManager{
		userManager: userManager,
	}
}

func (fm *FileManager) CreateFile(username, folderName, fileName, description string) error {
	user, err := fm.userManager.GetUser(username)
	if err != nil {
		return err
	}

	folder, exists := user.Folders[folderName]
	if !exists {
		return fmt.Errorf("folder '%s' does not exist for user '%s'", folderName, username)
	}

	if _, exists := folder.Files[fileName]; exists {
		return fmt.Errorf("file '%s' already exists in folder '%s' for user '%s'", fileName, folderName, username)
	}

	folder.Files[fileName] = &File{
		Name:        fileName,
		Description: description,
		CreatedAt:   time.Now(),
	}
	return nil
}

func (fm *FileManager) ListFiles(username, folderName, sortBy, order string) ([]*File, error) {
	user, err := fm.userManager.GetUser(username)
	if err != nil {
		return nil, err
	}

	folder, exists := user.Folders[folderName]
	if !exists {
		return nil, fmt.Errorf("folder '%s' does not exist for user '%s'", folderName, username)
	}

	files := make([]*File, 0, len(folder.Files))
	for _, file := range folder.Files {
		files = append(files, file)
	}

	switch sortBy {
	case "name":
		sort.Slice(files, func(i, j int) bool {
			if order == "asc" {
				return files[i].Name < files[j].Name
			}
			return files[i].Name > files[j].Name
		})
	case "created":
		sort.Slice(files, func(i, j int) bool {
			if order == "asc" {
				return files[i].CreatedAt.Before(files[j].CreatedAt)
			}
			return files[i].CreatedAt.After(files[j].CreatedAt)
		})
	}

	return files, nil
}
