package conv

import (
	"fmt"
	"reflect"
)

func ToString(i interface{}) (string, error) {
	i = indirectToStringerOrError(i)
	if i == nil {
		return "", fmt.Errorf("cannot convert nil to string")
	}
	switch v := i.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case fmt.Stringer:
		return v.String(), nil
	case error:
		return v.Error(), nil
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Bool, reflect.Float32, reflect.Float64:
		return fmt.Sprint(v.Interface()), nil
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return string(v.Bytes()), nil
		}
	}
	return "", fmt.Errorf("cannot convert %v to string", i)
}

func ToStringSlice(i interface{}) ([]string, error) {
	i = indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]string); ok {
		return l, nil
	}
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("cannot convert %v to slice", v.Kind())
	}
	num := v.Len()
	res := make([]string, num)
	var err error
	for j := 0; j < num; j++ {
		res[j], err = ToString(v.Index(j))
		if err != nil {
			return nil, fmt.Errorf("convert index %d: %w", i, err)
		}
	}
	return res, nil
}
