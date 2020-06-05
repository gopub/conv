package conv

import (
	"container/list"
	"reflect"
)

func ToList(i interface{}) *list.List {
	if i == nil {
		return list.New()
	}

	if l, ok := i.(*list.List); ok {
		return l
	}

	lt := reflect.TypeOf((*list.List)(nil))
	if it := reflect.TypeOf(i); it.ConvertibleTo(lt) {
		return reflect.ValueOf(i).Convert(lt).Interface().(*list.List)
	}

	l := list.New()
	v := reflect.ValueOf(indirect(i))
	if v.IsValid() && (v.Kind() == reflect.Slice || v.Kind() == reflect.Array) && !v.IsNil() {
		for j := 0; j < v.Len(); j++ {
			l.PushBack(v.Index(j).Interface())
		}
	} else {
		l.PushBack(i)
	}
	return l
}
