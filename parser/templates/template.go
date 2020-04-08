package templates

import (
	"go/token"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/RossMerr/jsonschema/parser"
)

var SchemaFuncMap = template.FuncMap{
	"mixedCase":   MixedCase,
	"typename":    Typename,
	"title":       strings.Title,
	"isStruct":    IsStruct,
	"isArray":     IsArray,
	"isString":    IsString,
	"isNumber":    IsNumber,
	"isInteger":   IsInteger,
	"isInterface": IsInterface,
	"isBoolean":   IsBoolean,
	"isReference": IsReference,
	"isEnum":      IsEnum,
	"isConst":     IsConst,
	"isMethod":    IsMethod,
	"isAllOf":     IsAllOf,
	"isAnyOf":     IsAnyOf,
	"isOneOf":     IsOneOf,
	"isType":      IsType,
}

func DefaultSchemaTemplate() (*template.Template, error) {
	tmpl, err := template.New("document").Funcs(SchemaFuncMap).Parse(DocumentTemplate)
	if err != nil {
		return nil, err
	}
	return AddTemplates(tmpl)

}

var TemplateArray = []string{
	StructTemplate,
	ArrayTemplate,
	NumberTemplate,
	IntegerTemplate,
	StringTemplate,
	BooleanTemplate,
	InterfaceTemplate,
	EnumTemplate,
	ReferenceTemplate,
	KindTemplate,
	ConstTemplate,
	parser.MethodTemplate,
	AllOfTemplate,
	AnyOfTemplate,
	OneOfTemplate,
	parser.MethodSignatureTemplate,
	TypeTemplate,
}

func AddTemplates(tmpl *template.Template) (*template.Template, error) {
	var err error
	for _, template := range TemplateArray {
		tmpl, err = tmpl.Parse(template)
		if err != nil {
			return nil, err
		}
	}

	return tmpl, nil
}

func MixedCase(raw string) string {
	if len(raw) < 1 {
		return raw
	}
	s := strings.Title(raw)
	return strings.ToLower(s[0:1]) + s[1:]
}

func Typename(raw string) string {
	if len(raw) < 1 {
		return raw
	}

	name := strings.TrimSuffix(raw, filepath.Ext(raw))

	// Valid field names must start with a unicode letter
	if !unicode.IsLetter(rune(name[0])) {
		name = "No " + name
	}

	// Valid field names must not be a reserved word
	if token.IsKeyword(name) {
		name = "Key " + name
	}

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	return reg.ReplaceAllString(strings.Title(clean), "")
}


func IsType(obj interface{}) bool {
	_, ok := obj.(*Type)
	return ok
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

func IsEnum(obj interface{}) bool {
	_, ok := obj.(*Enum)
	return ok
}

func IsConst(obj interface{}) bool {
	_, ok := obj.(*Const)
	return ok
}

func IsMethod(obj interface{}) bool {
	_, ok := obj.(*parser.Method)
	return ok
}

func IsAllOf(obj interface{}) bool {
	_, ok := obj.(*AllOf)
	return ok
}

func IsAnyOf(obj interface{}) bool {
	_, ok := obj.(*AnyOf)
	return ok
}

func IsOneOf(obj interface{}) bool {
	_, ok := obj.(*OneOf)
	return ok
}
