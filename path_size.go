package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PB = 1024 * TB
	EB = 1024 * PB
)

func formatSize(size int64, human bool) string {
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

func getDirSize(path string, human, recursive, all bool, totalSize *int64) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return "", err
			}
			if all || !strings.HasPrefix(entry.Name(), ".") {
    			*totalSize += info.Size()
			}

		}
	}

	if recursive {
		for _, entry := range entries {
			isDir := entry.IsDir()
			name := entry.Name()
			if isDir && (all || !strings.HasPrefix(name, ".")) {
				childPath := filepath.Join(path, name)
				_, err := getDirSize(childPath, human, recursive, all, totalSize)
				if err != nil {
					return "", err
				}
			}
		}
	}
	return formatSize(*totalSize, human), nil
}


// GetPathSize возвращает размер файла или директории по указанному пути.
//
// Если path указывает на директорию и recursive=true, размер считается рекурсивно
// (включая вложенные директории). Если recursive=false, учитываются только файлы
// первого уровня.
//
// Если human=true, размер возвращается в человекочитаемом формате (KB, MB и т.д.).
// Если all=false, скрытые файлы и директории (начинающиеся с ".") не учитываются.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	var totalSize int64
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return getDirSize(path, human, recursive, all, &totalSize)
	}
	if strings.HasPrefix(fileInfo.Name(), ".") && !all {
		return formatSize(0, human), nil
	}
	totalSize = fileInfo.Size()
	return formatSize(totalSize, human), nil
}
