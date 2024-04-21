package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ListImageFiles 用于列出指定文件夹中的所有图片文件
func ListImageFiles(folderPath string) ([]string, error) {
	var imagePaths []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查文件是否是图片文件（可根据需要修改文件扩展名列表）
		if isImageFile(path) {
			imagePaths = append(imagePaths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return imagePaths, nil
}

// 判断文件是否是图片文件（根据文件扩展名来判断）
func isImageFile(filename string) bool {
	// 可以根据需要添加其他图片文件扩展名，例如 ".jpg", ".png", ".gif" 等
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}

	ext := strings.ToLower(filepath.Ext(filename))

	for _, imageExt := range imageExtensions {
		if ext == imageExt {
			return true
		}
	}

	return false
}

func WriteToFile(obj interface{}, path string) error {
	// 打开或创建 JSON 文件
	outputFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 使用 json.NewEncoder 将对象编码并写入文件
	encoder := json.NewEncoder(outputFile)
	if err := encoder.Encode(obj); err != nil {
		return err
	}

	fmt.Printf("已写入 %s\n", path)

	return nil
}
