package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 循環執行命令。
func main() {
	fmt.Println("Welcome VFS...")
	userManager := NewUserManager()
	folderManager := NewFolderManager(userManager)
	fileManager := NewFileManager(userManager)

	reader := bufio.NewReader(os.Stdin)

	for {
		//default 格式
		fmt.Print("Enter command: ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		command = strings.TrimSpace(command)
		parts := strings.Split(command, " ")
		if len(parts) < 2 {
			fmt.Println("Invalid command")
			continue
		}

		switch parts[0] {
		//新增使用者
		case "register":
			if len(parts) != 2 {
				fmt.Println("Invalid command")
				continue
			}
			username := parts[1]
			err := userManager.AddUser(username)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("User '%s' registered successfully\n", username)
			}
		// 新增資料夾
		case "create-folder":
			if len(parts) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			username := parts[1]
			folderName := parts[2]
			err := folderManager.CreateFolder(username, folderName)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Folder '%s' created successfully for user '%s'\n", folderName, username)
			}
		// 列出資料夾清單
		case "list-folders":
			if len(parts) < 2 {
				fmt.Println("Invalid command")
				continue
			}
			username := parts[1]
			sortBy := ""
			order := "asc"
			if len(parts) >= 4 && parts[2] == "--sort" {
				sortBy = parts[3]
				if len(parts) == 5 {
					order = parts[4]
				}
			}
			folders, err := folderManager.ListFolders(username, sortBy, order)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				for _, folder := range folders {
					fmt.Printf("Folder Name: %s, Created At: %s\n", folder.Name, folder.CreatedAt.Format("2006-01-02 15:04:05"))
				}
			}
		//新增檔案
		case "create-file":
			if len(parts) < 5 {
				fmt.Println("Invalid command")
				continue
			}
			username := parts[1]
			folderName := parts[2]
			fileName := parts[3]
			description := strings.Join(parts[4:], " ")
			err := fileManager.CreateFile(username, folderName, fileName, description)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("File '%s' created successfully in folder '%s' for user '%s'\n", fileName, folderName, username)
			}
		// 列出資料夾底下的檔案清單
		case "list-files":
			if len(parts) < 3 {
				fmt.Println("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
				continue
			}
			username := parts[1]
			folderName := parts[2]
			sortBy := ""
			order := "asc"
			if len(parts) >= 5 && parts[3] == "--sort" {
				if parts[4] == "name" {
					sortBy = "name"
				} else if parts[4] == "created" {
					sortBy = "created"
				} else {
					fmt.Println("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
					continue
				}
				if len(parts) == 6 {
					order = parts[5]
				}
			}
			files, err := fileManager.ListFiles(username, folderName, sortBy, order)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				for _, file := range files {
					fmt.Printf("File Name: %s, Created At: %s, Description: %s\n", file.Name, file.CreatedAt.Format("2006-01-02 15:04:05"), file.Description)
				}
			}
		default:
			fmt.Println("Unrecognized command:", parts[0])
		}
	}
}
