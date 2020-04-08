package jsonschema

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		a []string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "Does not contain",
			args: args{
				a: []string{"foo", "bar"},
				b: "hello world",
			},
			want: false,
		},
		{
			name: "Contains",
			args: args{
				a: []string{"foo", "bar", "hello world"},
				b: "hello world",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		a        []string
		delegate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Filter",
			args: args{
				a:        []string{"foo", "bar", "hello world"},
				delegate: func(s string) bool { return s == "foo" || s == "bar" },
			},
			want: []string{"foo", "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.a, tt.args.delegate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
