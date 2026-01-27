package code

import (
	"fmt"
	"os"
)

func getDirSize(path string) (string, error){
	entries, err := os.ReadDir(path)
	var totalSize int64 
	if err != nil {
		return "", err
	}
	for _, entry := range entries{
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return "", err
			}
			totalSize += info.Size()	
		}
	}
	return fmt.Sprintf("%d\t%s", totalSize, path), nil
}

func GetSize(path string) (string, error){
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return getDirSize(path)
	}
	return fmt.Sprintf("%d\t%s", fileInfo.Size(), path), nil
}