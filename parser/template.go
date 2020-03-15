package parser

import (
	"log"
	"text/template"
)

func Template() *template.Template {
	tmpl, err := template.New("document").Funcs(template.FuncMap{
		"toString":    ToString,
		"isStruct":    IsStruct,
		"isArray":     IsArray,
		"isString":    IsString,
		"isNumber":    IsNumber,
		"isInterface": IsInterface,
		"isBoolean":   IsBoolean,
	}).Parse(DocumentTemplate)
	tmpl, err = tmpl.Parse(StructTemplate)
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
	return tmpl
}
