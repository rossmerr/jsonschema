package templates

import (
	"strings"
	"text/template"

	"github.com/RossMerr/jsonschema/parser"
)

func Template() (*template.Template, error) {
	tmpl, err := template.New("document").Funcs(template.FuncMap{
		"isStruct":             IsStruct,
		"isArray":              IsArray,
		"isString":             IsString,
		"isNumber":             IsNumber,
		"isInteger":            IsInteger,
		"isInterface":          IsInterface,
		"isBoolean":            IsBoolean,
		"isReference":          IsReference,
		"isInterfaceReference": IsInterfaceReference,
		"isEnum":               IsEnum,
		"isConst":              IsConst,
		"isRoot":               IsRoot,
		"mixedCase":            MixedCase,
		"title":                strings.Title,
	}).Parse(parser.DocumentTemplate)
	tmpl, err = tmpl.Parse(StructTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(ArrayTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(NumberTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(IntegerTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(StringTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(BooleanTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(InterfaceTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(EnumTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(ReferenceTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(InterfaceReferenceTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(KindTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(ConstTemplate)
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(RootTemplate)
	if err != nil {
		return nil, err
	}
	return tmpl, nil

}

func IsStruct(obj interface{}) bool {
	_, ok := obj.(*Struct)
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

func IsInterfaceReference(obj interface{}) bool {
	_, ok := obj.(*InterfaceReference)
	return ok
}

func IsEnum(obj interface{}) bool {
	_, ok := obj.(*Enum)
	return ok
}

func IsConst(obj interface{}) bool {
	_, ok := obj.(*Const)
	return ok
}

func IsRoot(obj interface{}) bool {
	_, ok := obj.(*Root)
	return ok
}

func MixedCase(raw string) string {
	if len(raw) < 1 {
		return raw
	}
	s := strings.Title(raw)
	return strings.ToLower(s[0:1]) + s[1:]
}
