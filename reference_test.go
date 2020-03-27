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
			name:      "Pointer followed by fragments",
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

func TestReference_Base(t *testing.T) {
	tests := []struct {
		name string
		s    jsonschema.Reference
		want string
	}{
		{
			name: "Empty string",
			s:    jsonschema.NewReference(""),
			want: ".",
		},
		{
			name: "No pointer",
			s:    jsonschema.NewReference("test"),
			want: "test",
		},
		{
			name: "Just pointer",
			s:    jsonschema.NewReference("#"),
			want: ".",
		},
		{
			name: "One fragment",
			s:    jsonschema.NewReference("#test"),
			want: ".",
		},
		{
			name: "Two fragment",
			s:    jsonschema.NewReference("#hello/world"),
			want: ".",
		},
		{
			name: "Relative",
			s:    jsonschema.NewReference("test.json#hello/world"),
			want: "test.json",
		},
		{
			name: "Relative",
			s:    jsonschema.NewReference("http://www.sample.com/test.json#hello/world"),
			want: "test.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Base(); got != tt.want {
				t.Errorf("Base() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReference_Pointer(t *testing.T) {
	tests := []struct {
		name string
		s    jsonschema.Reference
		want jsonschema.Pointer
	}{
		{
			name: "Empty string",
			s:    jsonschema.NewReference(""),
			want: jsonschema.Pointer(""),
		},
		{
			name: "No pointer",
			s:    jsonschema.NewReference("test"),
			want: jsonschema.Pointer(""),
		},
		{
			name: "Just pointer",
			s:    jsonschema.NewReference("#"),
			want: jsonschema.Pointer(""),
		},
		{
			name: "Pointer",
			s:    jsonschema.NewReference("#test"),
			want: jsonschema.Pointer("test"),
		},
		{
			name: "One fragment no pointer",
			s:    jsonschema.NewReference("#/test"),
			want: jsonschema.Pointer(""),
		},
		{
			name: "Pointer followed by two fragment",
			s:    jsonschema.NewReference("#test/hello/world"),
			want: jsonschema.Pointer("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pointer(); got != tt.want {
				t.Errorf("Pointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
