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
		return false, fmt.Errorf("cannot convert nil to bool")
	case string:
		return strconv.ParseBool(v)
	}

	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.String:
		return strconv.ParseBool(v.String())
	}

	n, err := ToInt64(i)
	if err != nil {
		return false, fmt.Errorf("cannot convert %v to bool", i)
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
