package conv

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func ToBool(i interface{}) (bool, error) {
	i = indirect(i)
	switch v := i.(type) {
	case bool:
		return v, nil
	case nil:
		return false, errNilValue
	case string:
		return strconv.ParseBool(v)
	}

	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.String:
		return strconv.ParseBool(v.String())
	}

	n, err := parseInt64(i)
	if err != nil {
		return false, fmt.Errorf("cannot convert %#v of type %T to bool", i, i)
	}
	return n != 0, nil
}

func MustBool(i interface{}) bool {
	v, err := ToBool(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToBoolSlice(i interface{}) ([]bool, error) {
	i = indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]bool); ok {
		return l, nil
	}
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("cannot convert %#v of type %T to []bool", i, i)
	}
	num := v.Len()
	res := make([]bool, num)
	var err error
	for j := 0; j < num; j++ {
		res[j], err = ToBool(v.Index(j).Interface())
		if err != nil {
			return nil, fmt.Errorf("convert index %d: %w", i, err)
		}
	}
	return res, nil
}

func MustBoolSlice(i interface{}) []bool {
	v, err := ToBoolSlice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}
