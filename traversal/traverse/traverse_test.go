package traverse

import (
	"reflect"
	"strings"
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestWalk(t *testing.T) {
	type args struct {
		s    *jsonschema.Schema
		path jsonschema.Path
	}
	tests := []struct {
		name string
		args args
		want *jsonschema.Schema
	}{
		{
			name: "Empty",
			args: args{
				s:    &jsonschema.Schema{},
				path: jsonschema.Path{},
			},
			want: &jsonschema.Schema{},
		},
		{
			name: "Definitions",
			args: args{
				s: &jsonschema.Schema{
					Definitions: map[string]*jsonschema.Schema{
						"test": &jsonschema.Schema{ID: jsonschema.ID("test")},
					},
				},
				path: jsonschema.Path{"definitions", "test"},
			},
			want: &jsonschema.Schema{ID: jsonschema.ID("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Walk(tt.args.s, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Walk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	type args struct {
		a        []string
		delegate func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "To lowercase",
			args: args{
				a:        []string{"Foo", "Bar"},
				delegate: func(s string) string { return strings.ToLower(s) },
			},
			want: []string{"foo", "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ForEach(tt.args.a, tt.args.delegate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForEach() = %v, want %v", got, tt.want)
			}
		})
	}
}
