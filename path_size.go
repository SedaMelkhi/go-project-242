package code

import (
	"fmt"
	"os"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PB = 1024 * TB
	EB = 1024 * PB
)

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	type Value struct {
		name string
		num  float64
	}
	values := []Value{
		{
			name: "EB",
			num:  EB,
		},
		{
			name: "PB",
			num:  PB,
		},
		{
			name: "TB",
			num:  TB,
		},
		{
			name: "GB",
			num:  GB,
		},
		{
			name: "MB",
			num:  MB,
		},
		{
			name: "KB",
			num:  KB,
		},
	}
	for _, value := range values {
		if float64(size) >= float64(value.num) {
			return fmt.Sprintf("%.1f%s", float64(size)/value.num, value.name)
		}
	}
	return fmt.Sprintf("%dB", size)
}

func getDirSize(path string, human bool) (string, error) {
	entries, err := os.ReadDir(path)
	var totalSize int64
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return "", err
			}
			totalSize += info.Size()
		}
	}
	return fmt.Sprintf("%s\t%s", FormatSize(totalSize, human), path), nil
}

func GetSize(path string, human bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return getDirSize(path, human)
	}
	return fmt.Sprintf("%s\t%s", FormatSize(fileInfo.Size(), human), path), nil
}
