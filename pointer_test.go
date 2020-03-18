package jsonschema_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestPointer_Fragments(t *testing.T) {
	tests := []struct {
		name      string
		s         jsonschema.Pointer
		wantQuery []string
	}{
		{
			name:      "Empty string",
			s:         jsonschema.NewPointer(""),
			wantQuery: []string{},
		},
		{
			name:      "No pointer",
			s:         jsonschema.NewPointer("test"),
			wantQuery: []string{},
		},
		{
			name:      "Just pointer",
			s:         jsonschema.NewPointer("#"),
			wantQuery: []string{},
		},
		{
			name:      "One fragment",
			s:         jsonschema.NewPointer("#test"),
			wantQuery: []string{"test"},
		},
		{
			name:      "Two fragment",
			s:         jsonschema.NewPointer("#hello/world"),
			wantQuery: []string{"hello", "world"},
		},
		{
			name:      "Relative",
			s:         jsonschema.NewPointer("test.json#hello/world"),
			wantQuery: []string{"hello", "world"},
		},
		{
			name:      "Relative",
			s:         jsonschema.NewPointer("http://www.sample.com/test.json#hello/world"),
			wantQuery: []string{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotQuery := tt.s.Fragments(); !reflect.DeepEqual(gotQuery, tt.wantQuery) {
				t.Errorf("Fragments() = %v, want %v", gotQuery, tt.wantQuery)
			}
		})
	}
}

func TestPointer_Base(t *testing.T) {
	tests := []struct {
		name string
		s    jsonschema.Pointer
		want string
	}{
		{
			name: "Empty string",
			s:    jsonschema.NewPointer(""),
			want: ".",
		},
		{
			name: "No pointer",
			s:    jsonschema.NewPointer("test"),
			want: "test",
		},
		{
			name: "Just pointer",
			s:    jsonschema.NewPointer("#"),
			want: ".",
		},
		{
			name: "One fragment",
			s:    jsonschema.NewPointer("#test"),
			want: ".",
		},
		{
			name: "Two fragment",
			s:    jsonschema.NewPointer("#hello/world"),
			want: ".",
		},
		{
			name: "Relative",
			s:    jsonschema.NewPointer("test.json#hello/world"),
			want: "test.json",
		},
		{
			name: "Relative",
			s:    jsonschema.NewPointer("http://www.sample.com/test.json#hello/world"),
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
