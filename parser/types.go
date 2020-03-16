package parser

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Types interface {
	Comment()     string
	ID() jsonschema.ID
}

func IsStruct(obj interface{}) bool {
	_, ok := obj.(*AnonymousStruct)
	return ok
}

func IsInterface(obj interface{}) bool {
	_, ok := obj.(*Interface)
	return ok
}

func IsArray(obj interface{}) bool {
	_, ok := obj.(*Array)
	return ok
}

func IsString(obj interface{}) bool {
	_, ok := obj.(*String)
	return ok
}

func IsNumber(obj interface{}) bool {
	_, ok := obj.(*Number)
	return ok
}

func IsBoolean(obj interface{}) bool {
	_, ok := obj.(*Boolean)
	return ok
}

func ToString(raw json.RawMessage) string {
	var s string
	err := json.Unmarshal(raw, &s)
	if err != nil {
		panic(err)
	}
	return s
}

func MixedCase(raw string) string {
	if len(raw) < 1 {
		return raw
	}
	s := strings.Title(raw)
	return  strings.ToLower(s[0:1]) + s[1:]
}

func KindOf(src interface{}) string {
	return reflect.ValueOf(src).Kind().String()
}