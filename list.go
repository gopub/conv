package conv

import (
	"container/list"
	"reflect"
)

// ToList creates list.List
// i can be nil, *list.List, or array/slice
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
	v := reflect.ValueOf(Indirect(i))
	if v.IsValid() && (v.Kind() == reflect.Slice || v.Kind() == reflect.Array) && !v.IsNil() {
		for j := 0; j < v.Len(); j++ {
			l.PushBack(v.Index(j).Interface())
		}
	} else {
		l.PushBack(i)
	}
	return l
}
