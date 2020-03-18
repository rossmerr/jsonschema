package jsonschema

import (
	"testing"
)

// func TestNewID_Typename(t *testing.T) {
// 	type args struct {
// 		schema string
// 		path string
// 		typename string
// 		err error
// 	}
// 	tests := []struct {
// 		name string
// 		s string
// 		want args
// 	}{
// 		{
// 			name: "TimeType",
// 			s:"TimeType",
// 			want: args{
// 				schema: "",
// 				path: "",
// 				typename: "TimeType",
// 				err: nil,
// 			},
// 		},
// 		{
// 			name: "TimeType",
// 			s:"#/definitions/TimeType",
// 			want: args{
// 				schema: "",
// 				path: "/definitions/",
// 				typename: "TimeType",
// 				err: nil,
// 			},
// 		},
// 		{
// 			name: "TimeType",
// 			s:"measureable_units.json#/definitions/TimeType",
// 			want: args{
// 				schema: "measureable_units.json",
// 				path: "/definitions/",
// 				typename: "TimeType",
// 				err: nil,
// 			},
// 		},
// 		{
// 			name: "TimeType",
// 			s:"https://raw.githubusercontent.com/beerjson/beerjson/master/json/measureable_units.json#/definitions/TimeType",
// 				want: args{
// 					schema: "https://raw.githubusercontent.com/beerjson/beerjson/master/json/measureable_units.json",
// 					path: "/definitions/",
// 					typename: "TimeType",
// 					err: nil,
// 				},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if schema, path, typename, err := NewID(tt.s).Parts(); err != tt.want.err || schema != tt.want.schema || path != tt.want.path || typename != tt.want.typename {
// 				t.Errorf("NewID()")
// 			}
// 		})
// 	}
// }

func TestID_Pointer(t *testing.T) {
	tests := []struct {
		name     string
		s        ID
		wantPath string
		wantQuery []string
		wantError bool
	}{
		{
			name:"Empty string",
			s: ID(""),
			wantPath:"",
			wantQuery: nil,
			wantError: true,
		},
		{
			name:"No #",
			s: ID("test"),
			wantPath:"",
			wantQuery:nil,
			wantError: true,
		},
		{
			name:"Root ID",
			s: ID("#test"),
			wantPath:"",
			wantQuery:[]string{"test"},
			wantError: false,
		},
		{
			name:"Defintions",
			s: ID("#defintions/test"),
			wantPath:"",
			wantQuery:[]string{"defintions","test"},
			wantError: false,
		},
		{
			name:"Relative",
			s: ID("test.json#defintions/test"),
			wantPath:"test.json",
			wantQuery:[]string{"defintions","test"},
			wantError: false,
		},
		{
			name:"Absolute",
			s: ID("http://www.test.com/test.json#defintions/test"),
			wantPath:"http://www.test.com/test.json",
			wantQuery:[]string{"defintions","test"},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, gotPath, gotError := tt.s.Pointer()
			if (gotError != nil) != tt.wantError {
				t.Errorf("Pointer() gotError = %v, want %v", gotError, tt.wantError)
			}

			if gotFile != tt.wantPath {
				t.Errorf("Pointer() gotFile = %v, want %v", gotFile, tt.wantPath)
			}
			if !Equal(gotPath, tt.wantQuery)  {
				t.Errorf("Pointer() gotPath = %v, want %v", gotPath, tt.wantQuery)
			}
		})
	}
}

func TestID_IsAbs(t *testing.T) {
	tests := []struct {
		name string
		s    ID
		want bool
	}{
		{
			name:"Empty string",
			s: ID(""),
			want: false,
		},
		{
			name:"No #",
			s: ID("test"),
			want: false,
		},
		{
			name:"Root ID",
			s: ID("#test"),
			want: false,
		},
		{
			name:"Relative",
			s: ID("test.json#defintions/test"),
			want: false,
		},
		{
			name:"Absolute",
			s: ID("http://www.test.com/test.json#defintions/test"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsAbs(); got != tt.want {
				t.Errorf("IsAbs() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestID_Typename2(t *testing.T) {
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
			if got := tt.s.Typename(); got != tt.want {
				t.Errorf("Typename() = %v, want %v", got, tt.want)
			}
		})
	}
}