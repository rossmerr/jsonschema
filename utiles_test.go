package jsonschema

import (
	"reflect"
	"strings"
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

func TestKeysString(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Key List",
			args: args{m: map[string]string{"foo": "bar", "hello": "world"}},
			want: []string{"foo=bar,hello=world", "hello=world,foo=bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeysString(tt.args.m)
			if !(got == tt.want[0] || got == tt.want[1]) {
				t.Errorf("KeysString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "Unicode letters",
			args: args{s: "foo bar"},
			want: "FooBar",
		},
		{
			name: "Non alphanumeric",
			args: args{s: "hello/world"},
			want: "HelloWorld",
		},
		{
			name: "Spaces",
			args: args{s: " foo bar "},
			want: "FooBar",
		},
		{
			name: "Spaces with non alphanumeric",
			args: args{s: "@ foo$bar #"},
			want: "FooBar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := title(tt.args.s); got != tt.want {
				t.Errorf("title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToTypename(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{s: ""},
			want: ".",
		},
		{
			name: "Dot string",
			args: args{s: "."},
			want: ".",
		},
		{
			name: "Unicode letter",
			args: args{s: "foo bar"},
			want: "FooBar",
		},
		{
			name: "Starting with a none unicode letter",
			args: args{s: "1 foo bar"},
			want: "No1FooBar",
		},
		{
			name: "Testing dash",
			args: args{s: "foo-bar"},
			want: "FooBar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTypename(tt.args.s); got != tt.want {
				t.Errorf("Structname() = %v, want %v", got, tt.want)
			}
		})
	}
}
