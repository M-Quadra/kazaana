package kazaana

import (
	"reflect"
	"time"
)

// Error error info
type Error struct {
	Raw error

	Time    time.Time
	Args    []any
	Callers []string
}

var ( // check interface
	_ Unwrapable = Error{}
	_ error      = Error{}
)

func (e Error) Error() string {
	err := e.Raw
	if err == nil {
		return ""
	}

	return e.Raw.Error()
}

func (e Error) Unwrap() error {
	return e.Raw
}

// As support errors.As to self
func (e Error) As(target any) bool {
	if e.Raw == nil || target == nil {
		return false
	}

	val := reflect.ValueOf(target)
	typ := val.Type()
	if typ.Kind() != reflect.Ptr {
		return false
	}

	targetType := typ.Elem()
	if reflect.TypeOf(&e).AssignableTo(targetType) {
		val.Elem().Set(reflect.ValueOf(&e))
		return true
	}
	return false
}
