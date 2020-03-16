package jsonschema

import "testing"

func TestNewID_Typename(t *testing.T) {
	type args struct {
		schema string
		path string
		typename string
		err error
	}
	tests := []struct {
		name string
		s string
		want args
	}{
		{
			name: "TimeType",
			s:"TimeType",
			want: args{
				schema: "",
				path: "",
				typename: "TimeType",
				err: nil,
			},
		},
		{
			name: "TimeType",
			s:"#/definitions/TimeType",
			want: args{
				schema: "",
				path: "/definitions/",
				typename: "TimeType",
				err: nil,
			},
		},
		{
			name: "TimeType",
			s:"measureable_units.json#/definitions/TimeType",
			want: args{
				schema: "measureable_units.json",
				path: "/definitions/",
				typename: "TimeType",
				err: nil,
			},
		},
		{
			name: "TimeType",
			s:"https://raw.githubusercontent.com/beerjson/beerjson/master/json/measureable_units.json#/definitions/TimeType",
				want: args{
					schema: "https://raw.githubusercontent.com/beerjson/beerjson/master/json/measureable_units.json",
					path: "/definitions/",
					typename: "TimeType",
					err: nil,
				},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if schema, path, typename, err := NewID(tt.s).Parts(); err != tt.want.err || schema != tt.want.schema || path != tt.want.path || typename != tt.want.typename {
				t.Errorf("NewID()")
			}
		})
	}
}

