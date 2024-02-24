package wd_mock_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

const (
	keyEnvDebug  = "CI_DEBUG"
	keyEnvCiNum  = "CI_NUMBER"
	keyEnvCiKey  = "CI_KEY"
	keyEnvCiKeys = "CI_KEYS"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	envDebug           = false

	envCiNum  = 0
	envCiKey  = ""
	envCiKeys []string
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()

	envDebug = fetchOsEnvBool(keyEnvDebug, false)
	envCiNum = fetchOsEnvInt(keyEnvCiNum, 1)
	envCiKey = fetchOsEnvStr(keyEnvCiKey, "")
	envCiKeys = fetchOsEnvArray(keyEnvCiKeys)
}

// test case basic tools start

// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
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

// pathIsDir
//
//	path is dir
func pathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// rmDir
//
//	remove dir by path
//
//nolint:golint,unused
func rmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := pathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// mkdir
//
//	will use FileMode 0766
func mkdir(path string) error {
	err := os.MkdirAll(path, os.FileMode(0766))
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

// readFileAsByte
//
//	read file as byte array
func readFileAsByte(path string) ([]byte, error) {
	exists, err := pathExists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("path not exist %v", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("path: %s read err: %v", path, err)
	}

	return data, nil
}

// readFileAsJson
//
//	read file as json
func readFileAsJson(path string, v interface{}) error {
	fileAsByte, errRead := readFileAsByte(path)
	if errRead != nil {
		return fmt.Errorf("path: %s , read file as err: %v", path, errRead)
	}
	err := json.Unmarshal(fileAsByte, v)
	if err != nil {
		return fmt.Errorf("path: %s , read file as json err: %v", path, err)
	}
	return nil
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

// writeFileAsJson write json file
//
//	path most use Abs Path
//	v data
//	fileMod os.FileMode(0666) or os.FileMode(0644)
//	coverage true will coverage old
//	beauty will format json when write
func writeFileAsJson(path string, v interface{}, fileMod fs.FileMode, coverage, beauty bool) error {
	if !coverage {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if beauty {
		var str bytes.Buffer
		err := json.Indent(&str, data, "", "  ")
		if err != nil {
			return err
		}
		return writeFileByByte(path, str.Bytes(), fileMod, coverage)
	}
	return writeFileByByte(path, data, fileMod, coverage)
}

// writeFileAsJsonBeauty
//
//	write json file as 0766 and beauty
func writeFileAsJsonBeauty(path string, v interface{}, coverage bool) error {
	return writeFileAsJson(path, v, os.FileMode(0766), coverage, true)
}

var currentTestDataFolderAbsPath = ""

// getOrCreateTestDataFullPath
//
//	get or create test data full path will under this package testdata
//	this function will create dir for return full path
func getOrCreateTestDataFullPath(elem ...string) (string, error) {
	if elem == nil || len(elem) < 1 {
		return "", fmt.Errorf("must has one elem")
	}
	dataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(dataFolderFullPath, elem[0])
	if len(elem) > 1 {
		for i := 1; i < len(elem); i++ {
			fullPath = filepath.Join(fullPath, elem[i])
		}
	}
	baseDir := filepath.Dir(fullPath)
	if !pathExistsFast(baseDir) {
		errMkdir := mkdir(baseDir)
		if errMkdir != nil {
			return fullPath, errMkdir
		}
	}
	if !pathIsDir(baseDir) {
		return "", fmt.Errorf("getOrCreateTestDataFullPath exist file, and can not create dir at path: %s", baseDir)
	}

	return fullPath, nil
}

// getOrCreateTestDataFolderFullPath
//
//	will create testdata folder under this package
func getOrCreateTestDataFolderFullPath() (string, error) {
	if currentTestDataFolderAbsPath != "" {
		return currentTestDataFolderAbsPath, nil
	}
	if testBaseFolderPath == "" {
		return "", fmt.Errorf("please testBaseFolderPath at each test case at init()")
	}
	currentTestDataFolderAbsPath = filepath.Join(testBaseFolderPath, "testdata")
	if !pathExistsFast(currentTestDataFolderAbsPath) {
		err := mkdir(currentTestDataFolderAbsPath)
		if err != nil {
			currentTestDataFolderAbsPath = ""
			return "", err
		}
	}
	return currentTestDataFolderAbsPath, nil
}

// fetchOsEnvBool
//
//	fetch os env by key.
//	if not found will return devValue.
//	return env not same as true (will be lowercase, so TRUE is same)
func fetchOsEnvBool(key string, devValue bool) bool {
	if os.Getenv(key) == "" {
		return devValue
	}
	return strings.ToLower(os.Getenv(key)) == "true"
}

// fetchOsEnvInt
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to int, return devValue
func fetchOsEnvInt(key string, devValue int) int {
	if os.Getenv(key) == "" {
		return devValue
	}
	outNum, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return devValue
	}

	return outNum
}

// fetchOsEnvStr
//
//	fetch os env by key.
//	return not found will return devValue.
func fetchOsEnvStr(key, devValue string) string {
	if os.Getenv(key) == "" {
		return devValue
	}
	return os.Getenv(key)
}

// fetchOsEnvInt
//
//	fetch os env split by `,` and trim space
//	return not found will return []string(nil).
func fetchOsEnvArray(key string) []string {
	var devValueStr []string
	if os.Getenv(key) == "" {
		return devValueStr
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	if len(splitVal) == 0 {
		return devValueStr
	}
	for _, item := range splitVal {
		devValueStr = append(devValueStr, strings.TrimSpace(item))
	}

	return devValueStr
}

// setEnvStr
//
//	set env by key and val
func setEnvStr(t *testing.T, key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		t.Fatalf("set env key [%v] string err: %v", key, err)
	}
}

// setEnvBool
//
//	set env by key and val
//
//nolint:golint,unused
func setEnvBool(t *testing.T, key string, val bool) {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		t.Fatalf("set env key [%v] bool err: %v", key, err)
	}
}

// setEnvU64
//
//	set env by key and val
//
//nolint:golint,unused
func setEnvU64(t *testing.T, key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		t.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// setEnvInt64
//
//	set env by key and val
//
//nolint:golint,unused
func setEnvInt64(t *testing.T, key string, val int64) {
	err := os.Setenv(key, strconv.FormatInt(val, 10))
	if err != nil {
		t.Fatalf("set env key [%v] int64 err: %v", key, err)
	}
}

// findAllEnvByPrefix4Print
//
//	find out all env by prefix for print
//	if not found will return ""
//
//nolint:golint,unused
func findAllEnvByPrefix4Print(prefix string) string {
	var sb strings.Builder
	for _, e := range os.Environ() {
		if strings.Index(e, prefix) == 0 {
			sb.WriteString(e)
			sb.WriteString("\n")
		}
	}
	if sb.Len() > 0 {
		return sb.String()
	}
	return ""
}

// findAllEnvByPrefix
//
//	find out all env by prefix
//	if not found will return nil
//
//nolint:golint,unused
func findAllEnvByPrefix(prefix string) map[string]string {
	var out map[string]string
	for _, e := range os.Environ() {
		if strings.Index(e, prefix) == 0 {
			envSplit := strings.Split(e, "=")
			if len(envSplit) > 1 {
				if out == nil {
					out = make(map[string]string)
				}
				out[envSplit[0]] = envSplit[1]
			} else if len(envSplit) > 0 {
				if out == nil {
					out = make(map[string]string)
				}
				out[envSplit[0]] = ""
			} else {
				continue
			}
		}
	}
	if out != nil && len(out) > 0 {
		return out
	}
	return nil
}

// test case basic tools end
