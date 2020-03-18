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
		"isInterface": IsInterface,
		"isBoolean":   IsBoolean,
		"mixedCase": MixedCase,
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
	return tmpl, nil
}
