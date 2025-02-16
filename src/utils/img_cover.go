package utils

import (
	"blog1/src/model"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func SaveLocalFile(file io.Reader, fileHeader *multipart.FileHeader) (filePath string, err error) {
	filePath = getFilePath(fileHeader)

	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ioutil.ReadAll err: ", err)
		return
	}

	// 写入文件
	err = ioutil.WriteFile(filePath, b, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil.WriteFile() err: ", err)
		return
	}

	return
}

func getFilePath(fileHeader *multipart.FileHeader) (filePath string) {
	// 使用相对路径
	filePath = filepath.Join(model.ReferenceDir, fileHeader.Filename)

	// 用 '/' 替换 '\'
	filePath = strings.Replace(filePath, "\\", "/", -1)

	return
}
