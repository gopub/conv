package conv

import (
	"log"
	"reflect"
	"runtime"
)

func FuncNameOf(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func NameOf(i interface{}) string {
	return reflect.TypeOf(i).Name()
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func ReverseArray(a interface{}) {
	a = Indirect(a)
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		log.Panicf("Cannot revert %T", a)
	}

	for i, j := 0, v.Len()-1; i < j; i, j = i+1, j-1 {
		vi, vj := v.Index(i), v.Index(j)
		tmp := vi.Interface()
		vi.Set(vj)
		vj.Set(reflect.ValueOf(tmp))
	}
}
