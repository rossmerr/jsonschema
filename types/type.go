package types

import (
	"reflect"
)

type Types string

type Type interface {
	Type() reflect.Kind
}