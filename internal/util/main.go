package util

import (
	"fmt"
	"os"
)

func CreateFolder(dirPath string) error {
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
	}
	return nil
}

func CreateFile(dirPath string, name string) error {
	filePath := dirPath + name
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	return nil
}

func DoesDirectoryExist(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func DoesFileExist(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ConvertMapKeysToSlice[T comparable](m map[string]T) []string {
	var keys []string

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}
