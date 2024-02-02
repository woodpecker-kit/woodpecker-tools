package env_transfer

import (
	"fmt"
	"github.com/joho/godotenv"
	"path/filepath"
)

func OverloadEnvFromFile(root string, fileName string) (map[string]string, error) {
	readerTarget := filepath.Join(root, fileName)
	if !pathExistsFast(readerTarget) {
		return nil, fmt.Errorf("OverloadEnvFromFile not exist file at path: %s", readerTarget)
	}
	envMap, errRead := godotenv.Read(readerTarget)
	if errRead != nil {
		return nil, fmt.Errorf("OverloadEnvFromFile godotenv.Read err: %v", errRead)
	}
	err := godotenv.Overload(readerTarget)
	if err != nil {
		return nil, fmt.Errorf("OverloadEnvFromFile godotenv.Overload err: %v", err)
	}
	return envMap, nil
}
