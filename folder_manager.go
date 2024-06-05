package main

import (
	"fmt"
	"sort"
	"time"
)

// 資料夾
type Folder struct {
	Name      string           // 資料夾名稱
	CreatedAt time.Time        // 新增時間
	Files     map[string]*File // 資料夾內的檔案
}

type FolderManager struct {
	userManager *UserManager
}

func NewFolderManager(userManager *UserManager) *FolderManager {
	return &FolderManager{
		userManager: userManager,
	}
}

// CreateFolder 新增資料夾
func (fm *FolderManager) CreateFolder(username, folderName string) error {
	user, err := fm.userManager.GetUser(username)
	if err != nil {
		return err
	}
	if _, exists := user.Folders[folderName]; exists {
		return fmt.Errorf("folder '%s' already exists for user '%s'", folderName, username)
	}
	user.Folders[folderName] = &Folder{
		Name:      folderName,
		CreatedAt: time.Now(),
		Files:     make(map[string]*File),
	}
	return nil
}

// ListFolders 資料夾清單
func (fm *FolderManager) ListFolders(username, sortBy, order string) ([]*Folder, error) {
	user, err := fm.userManager.GetUser(username)
	if err != nil {
		return nil, err
	}

	folders := make([]*Folder, 0, len(user.Folders))
	for _, folder := range user.Folders {
		folders = append(folders, folder)
	}

	switch sortBy {
	case "name":
		sort.Slice(folders, func(i, j int) bool {
			if order == "asc" {
				return folders[i].Name < folders[j].Name
			}
			return folders[i].Name > folders[j].Name
		})
	case "created":
		sort.Slice(folders, func(i, j int) bool {
			if order == "asc" {
				return folders[i].CreatedAt.Before(folders[j].CreatedAt)
			}
			return folders[i].CreatedAt.After(folders[j].CreatedAt)
		})
	}

	return folders, nil
}
