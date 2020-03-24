package parser

import (
	"strings"
)

type Types interface {
	Comment()     string
	ID() string
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

func IsInteger(obj interface{}) bool {
	_, ok := obj.(*Integer)
	return ok
}

func IsBoolean(obj interface{}) bool {
	_, ok := obj.(*Boolean)
	return ok
}

func IsReference(obj interface{}) bool {
	_, ok := obj.(*Reference)
	return ok
}

func IsEmbeddedStruct(obj interface{}) bool {
	_, ok := obj.(*EmbeddedStruct)
	return ok
}

func IsInterfaceReference(obj interface{}) bool {
	_, ok := obj.(*InterfaceReference)
	return ok
}


func MixedCase(raw string) string {
	if len(raw) < 1 {
		return raw
	}
	s := strings.Title(raw)
	return  strings.ToLower(s[0:1]) + s[1:]
}
