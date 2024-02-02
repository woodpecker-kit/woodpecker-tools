package env_transfer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// WriteEnv2File
//
//	write env to file
//	envData map[string]string will sort by key
func WriteEnv2File(root, fileName string, envData map[string]string) (string, error) {
	if len(envData) == 0 {
		return "", fmt.Errorf("envData is empty")
	}
	var keySort []string

	for key := range envData {
		if key == "" {
			continue
		}
		keySort = append(keySort, key)
	}
	sort.Strings(keySort)

	targetPath := filepath.Join(root, fileName)
	var sb strings.Builder

	for _, k := range keySort {
		if k == "" {
			continue
		}
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(`"`)
		sb.WriteString(envData[k])
		sb.WriteString(`"`)
		sb.WriteString("\n")
	}

	errWrite := writeFileFast(targetPath, []byte(sb.String()), true)
	if errWrite != nil {
		return "", errWrite
	}

	return targetPath, nil
}

func writeFileFast(path string, data []byte, coverage bool) error {
	return writeFileByByte(path, data, os.FileMode(0766), coverage)
}

// writeFileByByte
//
//	write bytes to file
//	path most use Abs Path
//	data []byte
//	fileMod os.FileMode(0666) or os.FileMode(0644)
//	coverage true will coverage old
func writeFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !pathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, fileMod)
		if err != nil {
			return fmt.Errorf("can not WriteFileByByte at new dir at mode: %v , at parent path: %v", fileMod, parentPath)
		}
	}
	err := os.WriteFile(path, data, fileMod)
	if err != nil {
		return fmt.Errorf("write data at path: %v, err: %v", path, err)
	}
	return nil
}

// pathExists
//
//	path exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// pathExistsFast
//
//	path exists fast
func pathExistsFast(path string) bool {
	exists, _ := pathExists(path)
	return exists
}
