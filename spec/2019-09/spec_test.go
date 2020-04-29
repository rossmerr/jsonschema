package spec_2019_09_test

import (
	"encoding/json"
	"go/ast"
	"reflect"
	"testing"

	spec_2019_09 "github.com/RossMerr/jsonschema/spec/2019-09"
)

func TestKeyword(t *testing.T) {
	type args struct {
		key      string
		property json.RawMessage
	}
	tests := []struct {
		name    string
		args    args
		want    *ast.TypeSpec
		wantErr bool
	}{
		{
			name: "string",
			args: args{
				key: "$comment",
				property: []byte(`{"type": "string" }`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := spec_2019_09.Keyword(tt.args.key, tt.args.property)

//			file, err := parser.ParseFile(fset, "demo", src, parser.ParseComments)

			if (err != nil) != tt.wantErr {
				t.Errorf("Keyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keyword() got = %v, want %v", got, tt.want)
			}
		})
	}
}