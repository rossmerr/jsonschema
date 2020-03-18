package parser

import (
	"log"
	"text/template"
)


func TemplateStruct() *template.Template {
	tmpl, err := template.New("document").Funcs(template.FuncMap{
		"toString":    ToString,
		"isStruct":    IsStruct,
		"isArray":     IsArray,
		"isString":    IsString,
		"isNumber":    IsNumber,
		"isInterface": IsInterface,
		"isBoolean":   IsBoolean,
		"kindOf": KindOf,
		"mixedCase": MixedCase,
	}).Parse(StructTemplate)
	tmpl, err = tmpl.Parse(AnonymousStructTemplate)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.Parse(ArrayTemplate)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.Parse(NumberTemplate)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.Parse(StringTemplate)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.Parse(BooleanTemplate)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.Parse(InterfaceTemplate)
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}
