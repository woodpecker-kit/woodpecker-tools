package env_mock

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"reflect"
	"strconv"
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
					env_kit.SetEnvBool(envTagVal, fVal.Bool())
				case reflect.String:
					vStr := fVal.String()
					if vStr == "" {
						envTagDefault, okMockEnvDefault := fType.Tag.Lookup("mock_env_default")
						if okMockEnvDefault {
							vStr = envTagDefault
						}
					}
					env_kit.SetEnvStr(envTagVal, vStr)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					env_kit.SetEnvU64(envTagVal, fVal.Uint())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					env_kit.SetEnvStr(envTagVal, strconv.FormatInt(fVal.Int(), 10))
				case reflect.Float32, reflect.Float64:
					env_kit.SetEnvStr(envTagVal, strconv.FormatFloat(fVal.Float(), 'f', -1, 64))
				}
			}
		}
	}
}
