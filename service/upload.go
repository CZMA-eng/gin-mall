package service

import (
	"gin_mall_tmp/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvartarToLocalStatic(file multipart.File, userId uint, username string)(filePath string, err error){
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath){
		CreateDir(basePath)
	}
	avatarPath := basePath + username + ".jpg" //TODO: 提取file后缀
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user"+bId+"/"+username+".jpg", nil
}


func DirExistOrNot(fileAddr string)bool{
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 
func CreateDir(dirName string) bool{
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}