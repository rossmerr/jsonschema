package jsonschema_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestAnchor_Fieldname(t *testing.T) {
	tests := []struct {
		name string
		s    jsonschema.Anchor
		want string
	}{
		{
			name: "Empty string",
			s:    jsonschema.NewAnchor(""),
			want: "",
		},
		{
			name: "No pointer",
			s:    jsonschema.NewAnchor("test"),
			want: "",
		},
		{
			name: "Just pointer",
			s:    jsonschema.NewAnchor("#"),
			want: "",
		},
		{
			name: "One fragment",
			s:    jsonschema.NewAnchor("#test"),
			want: "Test",
		},
		{
			name: "Two fragment",
			s:    jsonschema.NewAnchor("#hello/world"),
			want: "HelloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotQuery := tt.s.Fieldname(); !reflect.DeepEqual(gotQuery, tt.want) {
				t.Errorf("Path() = %v, want %v", gotQuery, tt.want)
			}
		})
	}
}
