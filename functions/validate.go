package functions

import "reflect"

func IsRequired(key string, required []string) bool {
	for _, s := range required {
		if key == s {
			return true
		}
	}

	return false
}

func IsNotRequired(key string, required []string) bool {
	return !IsRequired(key, required)
}

func Validate(dict map[string]interface{}) map[string]interface{} {
	for key, value :=  range dict {
		if key == "required" && value == false{
			delete(dict, key)
		}
		if isNil(value) {
			delete(dict, key)
		}
	}
	return dict
}

func isNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}