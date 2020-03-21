package conv

import (
	"fmt"
	"reflect"
)

func ToBytes(i interface{}) ([]byte, error) {
	i = indirect(i)
	switch v := i.(type) {
	case []byte:
		return v, nil
	case nil:
		return nil, errNilValue
	case string:
		return []byte(v), nil
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Slice && v.Type().Elem().Kind() == reflect.Uint8 {
		return v.Bytes(), nil
	}
	return nil, fmt.Errorf("cannot convert %v", v.Kind())
}
