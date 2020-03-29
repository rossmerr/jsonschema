package jsonschema_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestReference_Fragments(t *testing.T) {
	tests := []struct {
		name      string
		s         jsonschema.Reference
		wantQuery jsonschema.Path
	}{
		{
			name:      "Empty string",
			s:         jsonschema.NewReference(""),
			wantQuery: jsonschema.Path{},
		},
		{
			name:      "No pointer",
			s:         jsonschema.NewReference("test"),
			wantQuery: jsonschema.Path{},
		},
		{
			name:      "Just pointer",
			s:         jsonschema.NewReference("#"),
			wantQuery: jsonschema.Path{},
		},
		{
			name:      "No path in fragment",
			s:         jsonschema.NewReference("#test"),
			wantQuery: jsonschema.Path{},
		},
		{
			name:      "ID followed by fragments",
			s:         jsonschema.NewReference("#test/hello/world"),
			wantQuery: jsonschema.Path{"hello", "world"},
		},
		{
			name:      "Two fragment",
			s:         jsonschema.NewReference("#/hello/world"),
			wantQuery: jsonschema.Path{"hello", "world"},
		},
		{
			name:      "Relative",
			s:         jsonschema.NewReference("test.json#/hello/world"),
			wantQuery: jsonschema.Path{"hello", "world"},
		},
		{
			name:      "Relative",
			s:         jsonschema.NewReference("http://www.sample.com/test.json#/hello/world"),
			wantQuery: jsonschema.Path{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if path := tt.s.Path(); !reflect.DeepEqual(path, tt.wantQuery) {
				t.Errorf("Path() = %v, want %v", path, tt.wantQuery)
			}
		})
	}
}

func TestReference_Pointer(t *testing.T) {
	tests := []struct {
		name string
		s    jsonschema.Reference
		want jsonschema.ID
	}{
		{
			name: "Empty string",
			s:    jsonschema.NewReference(""),
			want: jsonschema.ID("."),
		},
		{
			name: "No pointer",
			s:    jsonschema.NewReference("test"),
			want: jsonschema.ID("."),
		},
		{
			name: "Just pointer",
			s:    jsonschema.NewReference("#"),
			want: jsonschema.ID("."),
		},
		{
			name: "ID",
			s:    jsonschema.NewReference("#test"),
			want: jsonschema.ID("."),
		},
		{
			name: "One fragment no pointer",
			s:    jsonschema.NewReference("#/test"),
			want: jsonschema.ID("."),
		},
		{
			name: "ID followed by two fragment",
			s:    jsonschema.NewReference("#test/hello/world"),
			want: jsonschema.ID("."),
		},
		{
			name: "ID followed by two fragment",
			s:    jsonschema.NewReference("https://www.sample.com/#test/hello/world"),
			want: jsonschema.ID("https://www.sample.com"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}
