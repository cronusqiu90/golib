package util

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		return true
	}
	return false
}

func WriteFile(path string, raw []byte) error {
	parentPath := filepath.Dir(path)
	if _, err := os.Stat(parentPath); os.IsNotExist(err) {
		if err := os.MkdirAll(parentPath, 0755); err != nil {
			return err
		}
	}
	return os.WriteFile(path, raw, 0755)
}

func ReadableSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	if size == 0 {
		return "0B"
	}
	exp := int(math.Log(float64(size)) / math.Log(1024))
	unit := units[exp]
	value := float64(size) / math.Pow(1024, float64(exp))
	formattedSize := fmt.Sprintf("%.2f%s", value, unit)
	return formattedSize
}
