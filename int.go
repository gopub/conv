package conv

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

const (
	MaxInt  = 1<<(8*unsafe.Sizeof(int(0))-1) - 1
	MinInt  = -1 << (8*unsafe.Sizeof(int(0)) - 1)
	MaxUint = 1<<(8*unsafe.Sizeof(uint(0))) - 1
)

func ToInt(i interface{}) (int, error) {
	n, err := ToInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to int", i)
	}
	if n > MaxInt || n < MinInt {
		return 0, fmt.Errorf("value %d out of range", n)
	}
	return int(n), nil
}

func ToInt8(i interface{}) (int8, error) {
	n, err := ToInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to int8", i)
	}
	if n > math.MaxInt8 {
		return 0, fmt.Errorf("value %d out of range", n)
	}
	return int8(n), nil
}

func ToInt16(i interface{}) (int16, error) {
	n, err := ToInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to int16", i)
	}
	if n > math.MaxInt16 {
		return 0, fmt.Errorf("value %d out of range", n)
	}
	return int16(n), nil
}

func ToInt32(i interface{}) (int32, error) {
	n, err := ToInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to int32", i)
	}
	if n > math.MaxInt32 {
		return 0, fmt.Errorf("value %d out of range", n)
	}
	return int32(n), nil
}

func ToInt64(i interface{}) (int64, error) {
	i = indirect(i)
	if i == nil {
		return 0, fmt.Errorf("cannot convert %v to int64", i)
	}
	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.Float32, reflect.Float64:
		return int64(v.Float()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return int64(v.Uint()), nil
	case reflect.Uint64:
		n := v.Uint()
		if n > math.MaxInt64 {
			return 0, fmt.Errorf("value %d out of range", n)
		}
		return int64(n), nil
	case reflect.String:
		n, err := strconv.ParseInt(v.String(), 0, 64)
		if err == nil {
			return n, nil
		}
		if errors.Is(err, strconv.ErrRange) {
			return 0, err
		}
		if f, fErr := strconv.ParseFloat(v.String(), 64); fErr == nil {
			return int64(f), nil
		}
		return 0, err
	default:
		return 0, fmt.Errorf("cannot convert %v to int64", i)
	}
}

func ToUint(i interface{}) (uint, error) {
	n, err := ToUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to uint", i)
	}
	if n > MaxUint {
		return 0, fmt.Errorf("value %d out of range", n)
	}
	return uint(n), nil
}

func ToUint8(i interface{}) (uint8, error) {
	n, err := ToUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to uint8", i)
	}
	if n > math.MaxUint8 {
		return 0, fmt.Errorf("value %d out of range", n)
	}
	return uint8(n), nil
}

func ToUint16(i interface{}) (uint16, error) {
	n, err := ToUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to uint16", i)
	}
	if n > math.MaxUint16 {
		return 0, fmt.Errorf("value %d out of range", n)
	}
	return uint16(n), nil
}

func ToUint32(i interface{}) (uint32, error) {
	n, err := ToUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %v to uint32", i)
	}
	if n > math.MaxUint32 {
		return 0, fmt.Errorf("value %d out of range", n)
	}
	return uint32(n), nil
}

func ToUint64(i interface{}) (uint64, error) {
	i = indirect(i)
	if i == nil {
		return 0, fmt.Errorf("cannot convert %v to uint64", i)
	}
	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		if f < 0 {
			return 0, fmt.Errorf("cannot convert %f to uint64", f)
		}
		return uint64(f), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := v.Int()
		if n < 0 {
			return 0, fmt.Errorf("cannot convert %d to uint64", n)
		}
		return uint64(n), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint(), nil
	case reflect.String:
		n, err := strconv.ParseInt(v.String(), 0, 64)
		if err == nil {
			if n < 0 {
				return 0, fmt.Errorf("cannot convert %d to uint64", n)
			}
			return uint64(n), nil
		}
		if errors.Is(err, strconv.ErrRange) {
			return 0, err
		}
		if f, fErr := strconv.ParseFloat(v.String(), 64); fErr == nil {
			if f < 0 {
				return 0, fmt.Errorf("cannot convert %f to uint64", f)
			}
			return uint64(f), nil
		}
		return 0, err
	default:
		return 0, fmt.Errorf("cannot convert %v to uint64", i)
	}
}
