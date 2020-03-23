package jsonschema

import (
	"testing"
)


func TestID_Base(t *testing.T) {
	tests := []struct {
		name    string
		s       ID
		want    string
	}{
		{
			name:"Empty string",
			s: ID(""),
			want: ".",
		},
		{
			name:"No #",
			s: ID("test"),
			want: ".",
		},
		{
			name:"Root ID",
			s: ID("#test"),
			want: ".",
		},
		{
			name:"Relative",
			s: ID("test.json#defintions/test"),
			want:"test.json",
		},
		{
			name:"Absolute",
			s: ID("http://www.test.com/test.json#defintions/test"),
			want:"test.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Base()

			if got != tt.want {
				t.Errorf("Base() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_Filename(t *testing.T) {
	tests := []struct {
		name string
		s    ID
		want string
	}{
		{
			name:"Empty string",
			s: ID(""),
			want:"",
		},
		{
			name:"No #",
			s: ID("test"),
			want:"",
		},
		{
			name:"Root ID",
			s: ID("#test"),
			want:"Test",
		},
		{
			name:"Defintions",
			s: ID("#defintions/test"),
			want:"Test",
		},
		{
			name:"Relative",
			s: ID("test.json#defintions/test"),
			want:"Test",
		},
		{
			name:"Absolute",
			s: ID("http://www.test.com/test.json#defintions/test"),
			want:"Test",
		},
		{
			name:"Test Case",
			s: ID("test.json#defintions/hello_world"),
			want:"HelloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Filename(); got != tt.want {
				t.Errorf("Fieldname() = %v, want %v", got, tt.want)
			}
		})
	}
}