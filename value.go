package rproxy

import (
	"reflect"
	"strconv"
)

type valueProxy struct {
	v reflect.Value
	p frame
	l string
}

var _ Proxy = (*valueProxy)(nil)

func (vp *valueProxy) Key(key interface{}) Proxy {
	k := vp.v.Kind()
	switch k {
	case reflect.Map:
		w := vp.v.MapIndex(reflect.ValueOf(key))
		if !w.IsValid() {
			return notFoundKey(vp, key)
		}
		return &valueProxy{
			v: deref(reflect.ValueOf(w)),
			p: vp,
			l: "." + w.String(),
		}
	case reflect.Struct:
		n := toStr(key)
		w := vp.v.FieldByName(n)
		if !w.IsValid() {
			return notFoundKey(vp, n)
		}
		return &valueProxy{
			v: deref(reflect.ValueOf(w)),
			p: vp,
			l: "." + n,
		}
	default:
		return typeError(vp, k, reflect.Map, reflect.Struct)
	}
}

func (vp *valueProxy) Index(i int) Proxy {
	k := vp.v.Kind()
	switch k {
	case reflect.Array, reflect.Slice:
		if i < 0 || i >= vp.v.Len() {
			return outOfIndex(vp, i)
		}
		w := vp.v.Index(i)
		return &valueProxy{
			v: deref(reflect.ValueOf(w)),
			p: vp,
			l: "[" + strconv.Itoa(i) + "]",
		}
	default:
		return typeError(vp, k, reflect.Array, reflect.Slice)
	}
}

func (vp *valueProxy) Value() (reflect.Value, error) {
	return vp.v, nil
}

func (vp *valueProxy) Interface() (interface{}, error) {
	if !vp.v.CanInterface() {
		return nil, forbidden(vp)
	}
	return vp.v.Interface(), nil
}

func (vp *valueProxy) Bool() (bool, error) {
	if k := vp.v.Kind(); k != reflect.Bool {
		return false, typeError(vp, k, reflect.Bool)
	}
	return vp.v.Bool(), nil
}

func (vp *valueProxy) Int() (int64, error) {
	k := vp.v.Kind()
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vp.v.Int(), nil
	default:
		return 0, typeError(vp, k, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64)
	}
}

func (vp *valueProxy) Uint() (uint64, error) {
	k := vp.v.Kind()
	switch k {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return vp.v.Uint(), nil
	default:
		return 0, typeError(vp, k, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64)
	}
}

func (vp *valueProxy) Float() (float64, error) {
	k := vp.v.Kind()
	switch k {
	case reflect.Float32, reflect.Float64:
		return vp.v.Float(), nil
	default:
		return 0, typeError(vp, k, reflect.Float32, reflect.Float64)
	}
}

func (vp *valueProxy) String() (string, error) {
	if k := vp.v.Kind(); k != reflect.String {
		return "", typeError(vp, k, reflect.String)
	}
	return vp.v.String(), nil
}

func (vp *valueProxy) parent() frame {
	return vp.p
}

func (vp *valueProxy) label() string {
	return vp.l
}
