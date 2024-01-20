package env_mock

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// MockEnvByStruct
//
//	 mock_env_key will mock env, only support string, bool, int, uint, float
//	 mock_env_default will mock env default value, only support string
//
//	example:
//			like:
//				type MockConfig struct {
//					KeyFoo string `mock_env_key:"PLUGIN_KEY_FOO" mock_env_default:"foo"`
//					KeyBar string `mock_env_key:"PLUGIN_KEY_BAR"`
//				}
//				type MockCase struct {
//					PluginDebug bool `mock_env_key:"PLUGIN_DEBUG"`
//					TimeOut     uint `mock_env_key:"PLUGIN_TIMEOUT"`
//					MockConfig MockConfig
//				}
func MockEnvByStruct(input interface{}) {
	value := reflect.ValueOf(input)

	typeInfo := value.Type()

	if typeInfo.Kind() == reflect.Ptr {
		typeInfo = typeInfo.Elem()
		if typeInfo.Kind() == reflect.Ptr {
			panic(fmt.Errorf("can not support double ptr, please check your input"))
		}
		value = value.Elem()
	}

	for i := 0; i < typeInfo.NumField(); i++ {
		// 取每个字段
		fVal := value.Field(i)
		fType := typeInfo.Field(i)

		if fType.Type.Kind() == reflect.Struct {
			MockEnvByStruct(fVal.Interface())
		} else {
			envTagVal, okEnvKey := fType.Tag.Lookup("mock_env_key")
			if okEnvKey {
				//fmt.Printf("tag [ mock_env_key : %s ]\n", envTagVal)

				switch fType.Type.Kind() {
				default:
					//fmt.Printf("Field [ %s ] : %v\n", fType.Name, fType.Type)
				case reflect.Bool:
					SetEnvBool(envTagVal, fVal.Bool())
				case reflect.String:
					vStr := fVal.String()
					if vStr == "" {
						envTagDefault, okMockEnvDefault := fType.Tag.Lookup("mock_env_default")
						if okMockEnvDefault {
							vStr = envTagDefault
						}
					}
					SetEnvStr(envTagVal, vStr)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					SetEnvU64(envTagVal, fVal.Uint())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					SetEnvStr(envTagVal, strconv.FormatInt(fVal.Int(), 10))
				case reflect.Float32, reflect.Float64:
					SetEnvStr(envTagVal, strconv.FormatFloat(fVal.Float(), 'f', -1, 64))
				}
			}
		}
	}
}

func SetEnvStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env key [%v] string err: %v", key, err)
	}
}

func SetEnvBool(key string, val bool) {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		log.Fatalf("set env key [%v] bool err: %v", key, err)
	}
}

func SetEnvU64(key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// FetchOsEnvBool
//
//	fetch os env by key.
//	if not found will return devValue.
//	return env not same as true (will be lowercase, so TRUE is same)
func FetchOsEnvBool(key string, devValue bool) bool {
	if os.Getenv(key) == "" {
		return devValue
	}
	return strings.ToLower(os.Getenv(key)) == "true"
}

// FetchOsEnvInt
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to int, return devValue
func FetchOsEnvInt(key string, devValue int) int {
	if os.Getenv(key) == "" {
		return devValue
	}
	outNum, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return devValue
	}

	return outNum
}

// FetchOsEnvStr
//
//	fetch os env by key.
//	return not found will return devValue.
func FetchOsEnvStr(key, devValue string) string {
	if os.Getenv(key) == "" {
		return devValue
	}
	return os.Getenv(key)
}

// FetchOsEnvArray
//
//	fetch os env split by `,` and trim space
//	return not found will return []string(nil).
func FetchOsEnvArray(key string) []string {
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
