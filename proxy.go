package rproxy

import (
	"fmt"
	"reflect"
)

// Proxy for reflect.Value
type Proxy interface {

	Value() (reflect.Value, error)

	Interface() (interface{}, error)

	Bool() (bool, error)

	Int() (int64, error)

	Uint() (uint64, error)

	Float() (float64, error)

	String() (string, error)

	// Key digs reflect.Value by key (wrap MapIndex or FieldByName).
	Key(key interface{}) Proxy

	// Index digs reflect.Value by index.
	Index(i int) Proxy

	frame
}

type frame interface {
	parent() frame
	label() string
}

// New creates a proxy for an interface{}.
func New(v interface{}) Proxy {
	return &valueProxy{v: deref(reflect.ValueOf(v))}
}

// deref returns reflect.Value with de-references
func deref(rv reflect.Value) reflect.Value {
	for {
		switch rv.Kind() {
		case reflect.Ptr, reflect.Interface:
			rv = rv.Elem()
		default:
			return rv
		}
	}
}

func toStr(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprint(v)
}
