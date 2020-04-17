package sample_schemas

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/RossMerr/jsonschema/interpreter"
	"github.com/RossMerr/jsonschema/sample_schemas/allOf"
	"github.com/RossMerr/jsonschema/sample_schemas/allOfObject"
	"github.com/RossMerr/jsonschema/sample_schemas/anyOf"
	"github.com/RossMerr/jsonschema/sample_schemas/basic"
	"github.com/RossMerr/jsonschema/sample_schemas/enum"
	"github.com/RossMerr/jsonschema/sample_schemas/nesting"
	"github.com/RossMerr/jsonschema/sample_schemas/oneOf"
	"github.com/RossMerr/jsonschema/sample_schemas/reference"
	"github.com/RossMerr/jsonschema/sample_schemas/referenceOutside"
)

func TestSchemas_Generate(t *testing.T) {
	// Redirect standard out to null, why because don't want fmt.Print outputs when running tests
	// stdout := os.Stdout
	// defer func() { os.Stdout = stdout }()
	// os.Stdout = os.NewFile(0, os.DevNull)

	type fields struct {
		paths       []string
		output      string
		sample      string
		str         interface{}
		packagename string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Basic",
			fields: fields{
				paths:       []string{"basic/schema.json"},
				packagename: "basic",
				output:      "basic/",
				sample:      "basic/sample.json",
				str: &struct {
					Person basic.Person `json:"person"`
				}{
					Person: basic.Person{},
				},
			},
		},
		{
			name: "Enum",
			fields: fields{
				paths:       []string{"enum/schema.json"},
				packagename: "enum",
				output:      "enum/",
				sample:      "enum/sample.json",
				str:         enum.KeyType(""),
			},
		},
		{
			name: "Nesting",
			fields: fields{
				paths:       []string{"nesting/schema.json"},
				packagename: "nesting",
				output:      "nesting/",
				sample:      "nesting/sample.json",
				str: &struct {
					Product nesting.Product `json:"product"`
				}{
					Product: nesting.Product{},
				},
			},
		},
		{
			name: "Reference",
			fields: fields{
				paths:       []string{"reference/schema.json"},
				packagename: "reference",
				output:      "reference/",
				sample:      "reference/sample.json",
				str: &struct {
					Home reference.HomeAddress `json:"home address"`
				}{
					Home: reference.HomeAddress{},
				},
			},
		},
		{
			name: "Reference Outside",
			fields: fields{
				paths:       []string{"referenceOutside/schema.json", "referenceOutside/outside.schema.json"},
				packagename: "referenceOutside",
				output:      "referenceOutside/",
				sample:      "referenceOutside/sample.json",
				str: &struct {
					Product referenceOutside.Product `json:"product"`
				}{
					Product: referenceOutside.Product{},
				},
			},
		},
		{
			name: "All Of",
			fields: fields{
				paths:       []string{"allOf/schema.json"},
				packagename: "allOf",
				output:      "allOf/",
				sample:      "allOf/sample.json",
				str: &struct {
					Product allOf.Allof `json:"allof"`
				}{
					Product: allOf.Allof{},
				},
			},
		},
		{
			name: "All Of Object",
			fields: fields{
				paths:       []string{"allOfObject/schema.json"},
				packagename: "allOfObject",
				output:      "allOfObject/",
				sample:      "allOfObject/sample.json",
				str: &struct {
					Product allOfObject.Allof `json:"allof"`
				}{
					Product: allOfObject.Allof{},
				},
			},
		},
		{
			name: "Any Of",
			fields: fields{
				paths:       []string{"anyOf/schema.json"},
				packagename: "anyOf",
				output:      "anyOf/",
				sample:      "anyOf/sample.json",
				str: &struct {
					Product anyOf.Anyof `json:"anyof"`
				}{
					Product: anyOf.Anyof{},
				},
			},
		},
		{
			name: "One of",
			fields: fields{
				paths:       []string{"oneOf/schema.json"},
				packagename: "oneOf",
				output:      "oneOf/",
				sample:      "oneOf/sample.json",
				str: &struct {
					Product oneOf.Oneof `json:"oneof"`
				}{
					Product: oneOf.Oneof{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			interpreter := interpreter.NewInterpreterDefaults(tt.fields.packagename)
			interpret, err := interpreter.Interpret(tt.fields.paths)
			if err != nil {
				t.Error(err)
			}
			_, err = interpret.ToFile(tt.fields.output)
			if err != nil {
				t.Error(err)
			}
			handler := func(w http.ResponseWriter, r *http.Request) {
				str := tt.fields.str

				defer r.Body.Close()

				err := json.NewDecoder(r.Body).Decode(&str)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				data, err := json.Marshal(str)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				io.WriteString(w, string(data))
			}

			file, err := ioutil.ReadFile(tt.fields.sample)
			if err != nil {
				t.Error(err)
			}

			reader := strings.NewReader(string(file))

			req := httptest.NewRequest("POST", "http://example.com", reader)
			w := httptest.NewRecorder()
			handler(w, req)

			err = ShouldEqualJSONObject(w.Body.Bytes(), file)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func ShouldEqualJSONObject(data1, data2 []byte) error {
	x := make(map[string]interface{})
	err := json.Unmarshal(data1, &x)
	if err != nil {
		return fmt.Errorf("unmarshal of data1 failed: %w", err)
	}
	y := make(map[string]interface{})
	err = json.Unmarshal(data2, &y)
	if err != nil {
		return fmt.Errorf("unmarshal of data2 failed: %w", err)
	}

	if !reflect.DeepEqual(x, y) {
		return fmt.Errorf("object not equal %v %v", x, y)
	}

	return nil
}
