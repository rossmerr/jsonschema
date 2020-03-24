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
		wantQuery []string
	}{
		{
			name:      "Empty string",
			s:         jsonschema.NewReference(""),
			wantQuery: []string{},
		},
		{
			name:      "No pointer",
			s:         jsonschema.NewReference("test"),
			wantQuery: []string{},
		},
		{
			name:      "Just pointer",
			s:         jsonschema.NewReference("#"),
			wantQuery: []string{},
		},
		{
			name:      "One fragment",
			s:         jsonschema.NewReference("#test"),
			wantQuery: []string{"test"},
		},
		{
			name:      "Two fragment",
			s:         jsonschema.NewReference("#hello/world"),
			wantQuery: []string{"hello", "world"},
		},
		{
			name:      "Relative",
			s:         jsonschema.NewReference("test.json#hello/world"),
			wantQuery: []string{"hello", "world"},
		},
		{
			name:      "Relative",
			s:         jsonschema.NewReference("http://www.sample.com/test.json#hello/world"),
			wantQuery: []string{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotQuery := tt.s.Pointer(); !reflect.DeepEqual(gotQuery, tt.wantQuery) {
				t.Errorf("Pointer() = %v, want %v", gotQuery, tt.wantQuery)
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
