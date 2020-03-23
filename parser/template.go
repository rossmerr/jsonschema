package parser

import (
	"text/template"
)

func Template() (*template.Template, error) {
	tmpl, err := template.New("document").Funcs(template.FuncMap{
		"isStruct":    IsStruct,
		"isArray":     IsArray,
		"isString":    IsString,
		"isNumber":    IsNumber,
		"isInteger":   IsInteger,
		"isInterface": IsInterface,
		"isBoolean":   IsBoolean,
		"isReference":   IsReference,
		"isEmbeddedStruct": IsEmbeddedStruct,
		"mixedCase":   MixedCase,
	}).Parse(StructTemplate)
	tmpl, err = tmpl.Parse(AnonymousStructTemplate)
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
	tmpl, err = tmpl.Parse(EmbeddedStructTemplate)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
