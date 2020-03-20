package conv

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func ToFloat32(i interface{}) (float32, error) {
	v, err := ToFloat64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to float32", i)
	}
	if v > math.MaxFloat32 || v < -math.MaxFloat32 {
		return 0, fmt.Errorf("value %f out of range", v)
	}
	return float32(v), nil
}

func ToFloat64(i interface{}) (float64, error) {
	i = indirect(i)
	if i == nil {
		return 0, fmt.Errorf("cannot convert %v to float64", i)
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return float64(v.Int()), nil
	case reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return float64(v.Uint()), nil
	case reflect.Uint64:
		return float64(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.String:
		return strconv.ParseFloat(v.String(), 64)
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %v to float64", i)
	}
}
