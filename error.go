package rproxy

import "reflect"

type errorProxy struct {
	p frame
	l string
	// TODO:
}

var (
	_ Proxy = (*errorProxy)(nil)
	_ error = (*errorProxy)(nil)
)

func (ep *errorProxy) Key(key interface{}) Proxy {
	return ep
}

func (ep *errorProxy) Index(i int) Proxy {
	return ep
}

func (ep *errorProxy) Value() (reflect.Value, error) {
	return reflect.Value{}, ep
}

func (ep *errorProxy) Interface() (interface{}, error) {
	return nil, ep
}

func (ep *errorProxy) Bool() (bool, error) {
	return false, ep
}

func (ep *errorProxy) Int() (int64, error) {
	return 0, ep
}

func (ep *errorProxy) Uint() (uint64, error) {
	return 0, ep
}

func (ep *errorProxy) Float() (float64, error) {
	return 0, ep
}

func (ep *errorProxy) String() (string, error) {
	return "", ep
}

func (ep *errorProxy) parent() frame {
	return ep.p
}

func (ep *errorProxy) label() string {
	return ep.l
}

// Error returns error message.
func (ep *errorProxy) Error() string {
	// TODO:
	return ""
}

// error constructors

func typeError(parent frame, actual, expected reflect.Kind, others ...reflect.Kind) *errorProxy {
	// TODO:
	return nil
}

func outOfIndex(parent frame, index int) *errorProxy {
	// TODO:
	return nil
}

func notFoundKey(parent frame, key interface{}) *errorProxy {
	// TODO:
	return nil
}

func forbidden(parent frame) *errorProxy {
	// TODO:
	return nil
}
