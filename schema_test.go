package jsonschema_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/interpreter"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/traversal"
)

func TestSchemas_Generate(t *testing.T) {
	type fields struct {
		documents map[jsonschema.ID]*jsonschema.Schema
		config    *jsonschema.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "Basic",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"basicBasic": loadRawSchema("samples/basic.json"),
				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output: "output/",
				},
			},
		},
		{
			name: "Nesting data structures",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"productNesting": loadRawSchema("samples/nesting.json"),
				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output:      "output/",
				},
			},
		},
		{
			name: "References inside the schema",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"http://example.com/reference.json": loadRawSchema("samples/reference.json"),
				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output:      "output/",
				},
			},
		},
		{
			name: "References outside the schema",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"https://example.com/reference-outside.schema.json": loadRawSchema("samples/reference-outside.schema.json"),
					"http://example.com/reference-outside.json": loadRawSchema("samples/reference-outside.json"),

				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output: "output/",
				},
			},
		},
		{
			name: "Oneof",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"http://example.com/oneof.json": loadRawSchema("samples/oneof.json"),
				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output:      "output/",
				},
			},
		},
		{
			name: "AnyOf",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"http://example.com/anyof.json": loadRawSchema("samples/anyof.json"),
				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output:      "output/",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := parser.NewParser(context.Background(), tt.fields.config.Packagename)
			parse := p.Parse(tt.fields.documents)
			interpret, err := interpreter.NewInterpretDefaults(parse)
			if err != nil {
				t.Error(err)
			}
			if err := interpret.ToFile(tt.fields.config.Output); (err != nil) != tt.wantErr {
				t.Errorf("Schemas.Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func loadRawSchema(filename string) *jsonschema.Schema {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var doc jsonschema.Schema
	err = json.Unmarshal(data, &doc)
	if err != nil {
		panic(err)
	}

	return &doc
}

func TestSchema_Traverse(t *testing.T) {
	type args struct {
		query []string
	}
	tests := []struct {
		name   string
		schema *jsonschema.Schema
		query []string
		want   *jsonschema.Schema
	}{
		{
			name: "Empty",
			schema: &jsonschema.Schema{},
			query:[]string{},
			want: nil,
		},
		{
			name: "Definitions",
			schema: &jsonschema.Schema{
				Definitions: map[string]*jsonschema.Schema{
					"test": &jsonschema.Schema{ID:jsonschema.ID("test")},
				},
			},
			query:[]string{"Definitions", "test"},
			want: &jsonschema.Schema{ID:jsonschema.ID("test")},
		},
		{
			name: "OneOf",
			schema: &jsonschema.Schema{
				OneOf:[]*jsonschema.Schema{
					&jsonschema.Schema{ID:jsonschema.ID("test")},
				},
			},
			query:[]string{"Oneof", "test"},
			want: &jsonschema.Schema{ID:jsonschema.ID("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traversal.Traverse(tt.schema, tt.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}