package spec_2019_09

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/RossMerr/jsonschema/spec/2019-09/keyword"
)

type Spec struct {
}

func Bootstrap() error {
	data, err := ioutil.ReadFile("schema.json")
	if err != nil {
		return err
	}

	x := make(map[string]json.RawMessage)
	err = json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	rawVocabulary := make(map[string]json.RawMessage)
	vocabulary := make(map[string]bool)
	if raw, ok := x[keyword.Vocabulary]; ok {
		err := json.Unmarshal(raw, &vocabulary)
		if err != nil {
			return err
		}

		for vocab, required := range vocabulary {
			if required == false {
				return fmt.Errorf("spec: the behavior of a false is undefined %v", vocab)
			}

			vocab = "meta/" + filepath.Base(vocab)
			data, err := ioutil.ReadFile(vocab)
			if err != nil {
				return fmt.Errorf("spec: %w", err)
			}

			m, err := RawProperties(data)
			if err != nil {
				return err
			}

			for key, prop := range m {
				if _, contains := rawVocabulary[key]; contains {
					return fmt.Errorf("spec: keyword %v already registered", key)
				}
				rawVocabulary[key] = prop
			}
		}
	}

	return nil
}

func RawProperties(data []byte) (map[string]json.RawMessage, error) {
	x := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &x)
	if err != nil {
		return nil, err
	}
	if raw, ok := x[keyword.Properties]; ok {
		properties := make(map[string]json.RawMessage)
		err = json.Unmarshal(raw, &properties)
		return properties, err
	}

	return nil, fmt.Errorf("spec: %v not found", keyword.Properties)
}

func Keywords(rawVocabulary map[string]json.RawMessage) ([]*ast.TypeSpec, error) {
	types := []*ast.TypeSpec{}

	for key, property := range rawVocabulary {

		spec, err := Keyword(key, property)
		if err != nil {
			return nil, err
		}

		types = append(types, spec)
	}

	return types, nil
}

func Keyword(key string, property json.RawMessage) (*ast.TypeSpec, error) {
	fieldList := &ast.FieldList{}
	spec := &ast.TypeSpec{
		Name: ast.NewIdent(Typename(key)),
		Type: &ast.StructType{
			Fields: fieldList,
		},
	}

	fields := make(map[string]json.RawMessage)
	err := json.Unmarshal(property, &fields)
	if err != nil {
		return nil, err
	}

	for field, t := range fields {

		field := &ast.Field{
			Names: []*ast.Ident{{Name: Typename(field)}},
			Type:  ast.NewIdent(string(t)),
		}

		fieldList.List = append(fieldList.List, field)
	}
	return spec, nil
}


func Typename(raw string) string {
	if len(raw) < 1 {
		return raw
	}

	name := strings.TrimSuffix(raw, filepath.Ext(raw))

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		panic(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	nospace :=  reg.ReplaceAllString(strings.Title(clean), "")

	// Valid field names must not be a reserved word
	if token.IsKeyword(nospace) {
		nospace = "Key " + nospace
	}

	return nospace
}
