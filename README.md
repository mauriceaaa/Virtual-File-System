# Virtual-File-System
## 功能

- 使用者註冊
- 建立資料夾
- 列出使用者的資料夾
- 建立檔案
- 列出資料夾中的檔案

### 命令列表

- `register [username]` - 使用者註冊
- `create-folder [username] [foldername]` - 新增資料夾
- `list-folders [username] [--sort [name|created] [asc|desc]]` - 列出使用者的資料夾，並可選排序方式
- `create-file [username] [foldername] [filename] [description]` - 在指定使用者的資料夾中創建檔案
- `list-files [username] [foldername] [--sort [name|created] [asc|desc]]` - 列出資料夾中的所有檔案，並可選排序方式

## 範例

```sh
register maurice
create-folder maurice helloworld
create-file maurice helloworld report.txt "Interview Report"
list-folders maurice --sort name asc
list-files maurice helloworld --sort created desc

## 文件結構

main.go - 主程式
user_manager.go - 使用者管理模組
folder_manager.go - 資料夾管理模組
file_manager.go - 檔案管理模組
user_manager_test.go - 使用者管理模組的單元測試
folder_manager_test.go - 資料夾管理模組的單元測試
file_manager_test.go - 檔案管理模組的單元測試

歡迎提交 PR 或 issue 來改進該項目。