package conv

import (
	"errors"
	"fmt"
	"log"
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
	n, err := parseInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to int64", i, i)
	}
	if n > MaxInt || n < MinInt {
		return 0, strconv.ErrRange
	}
	return int(n), nil
}

func MustInt(i interface{}) int {
	v, err := ToInt(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToInt8(i interface{}) (int8, error) {
	n, err := parseInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to int8", i, i)
	}
	if n > math.MaxInt8 || n < math.MinInt8 {
		return 0, strconv.ErrRange
	}
	return int8(n), nil
}

func MustInt8(i interface{}) int8 {
	v, err := ToInt8(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToInt16(i interface{}) (int16, error) {
	n, err := parseInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to int16", i, i)
	}
	if n > math.MaxInt16 || n < math.MinInt16 {
		return 0, strconv.ErrRange
	}
	return int16(n), nil
}

func MustInt16(i interface{}) int16 {
	v, err := ToInt16(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToInt32(i interface{}) (int32, error) {
	n, err := parseInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to int32", i, i)
	}
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0, strconv.ErrRange
	}
	return int32(n), nil
}

func MustInt32(i interface{}) int32 {
	v, err := ToInt32(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToInt64(i interface{}) (int64, error) {
	return parseInt64(i)
}

func MustInt64(i interface{}) int64 {
	v, err := parseInt64(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint(i interface{}) (uint, error) {
	n, err := parseUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to uint", i, i)
	}
	if n > MaxUint {
		return 0, strconv.ErrRange
	}
	return uint(n), nil
}

func MustUint(i interface{}) uint {
	v, err := ToUint(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint8(i interface{}) (uint8, error) {
	n, err := parseUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to uint8", i, i)
	}
	if n > math.MaxUint8 {
		return 0, strconv.ErrRange
	}
	return uint8(n), nil
}

func MustUint8(i interface{}) uint8 {
	v, err := ToUint8(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint16(i interface{}) (uint16, error) {
	n, err := parseUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to uint16", i, i)
	}
	if n > math.MaxUint16 {
		return 0, strconv.ErrRange
	}
	return uint16(n), nil
}

func MustUint16(i interface{}) uint16 {
	v, err := ToUint16(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint32(i interface{}) (uint32, error) {
	n, err := parseUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to uint32", i, i)
	}
	if n > math.MaxUint32 {
		return 0, strconv.ErrRange
	}
	return uint32(n), nil
}

func MustUint32(i interface{}) uint32 {
	v, err := ToUint32(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint64(i interface{}) (uint64, error) {
	return parseUint64(i)
}

func MustUint64(i interface{}) uint64 {
	v, err := parseUint64(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToIntSlice(i interface{}) ([]int, error) {
	i = indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]int); ok {
		return l, nil
	}
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("cannot convert %v to slice", v.Kind())
	}
	num := v.Len()
	res := make([]int, num)
	var err error
	for j := 0; j < num; j++ {
		res[j], err = ToInt(v.Index(j).Interface())
		if err != nil {
			return nil, fmt.Errorf("convert index %d: %w", j, err)
		}
	}
	return res, nil
}

func MustIntSlice(i interface{}) []int {
	v, err := ToIntSlice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToInt64Slice(i interface{}) ([]int64, error) {
	i = indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]int64); ok {
		return l, nil
	}
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("cannot convert %#v of type %T to []int64", i, i)
	}
	num := v.Len()
	res := make([]int64, num)
	var err error
	for j := 0; j < num; j++ {
		res[j], err = parseInt64(v.Index(j).Interface())
		if err != nil {
			return nil, fmt.Errorf("convert element at index %d: %w", i, err)
		}
	}
	return res, nil
}

func MustInt64Slice(i interface{}) []int64 {
	v, err := ToInt64Slice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUintSlice(i interface{}) ([]uint, error) {
	i = indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]uint); ok {
		return l, nil
	}
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("cannot convert %#v of type %T to []uint", i, i)
	}
	num := v.Len()
	res := make([]uint, num)
	var err error
	for j := 0; j < num; j++ {
		res[j], err = ToUint(v.Index(j).Interface())
		if err != nil {
			return nil, fmt.Errorf("convert element at index %d: %w", i, err)
		}
	}
	return res, nil
}

func MustUintSlice(i interface{}) []uint {
	v, err := ToUintSlice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint64Slice(i interface{}) ([]uint64, error) {
	i = indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]uint64); ok {
		return l, nil
	}
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("cannot convert %#v of type %T to []uint64", i, i)
	}
	num := v.Len()
	res := make([]uint64, num)
	var err error
	for j := 0; j < num; j++ {
		res[j], err = parseUint64(v.Index(j).Interface())
		if err != nil {
			return nil, fmt.Errorf("convert element at index %d: %w", i, err)
		}
	}
	return res, nil
}

func MustUint64Slice(i interface{}) []uint64 {
	v, err := ToUint64Slice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func parseInt64(i interface{}) (int64, error) {
	i = indirect(i)
	if i == nil {
		return 0, errNilValue
	}
	if b, ok := i.([]byte); ok {
		i = string(b)
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
			return 0, strconv.ErrRange
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
		return 0, strconv.ErrSyntax
	}
}

func parseUint64(i interface{}) (uint64, error) {
	i = indirect(i)
	if i == nil {
		return 0, errNilValue
	}
	if b, ok := i.([]byte); ok {
		i = string(b)
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
			return 0, strconv.ErrRange
		}
		return uint64(f), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := v.Int()
		if n < 0 {
			return 0, strconv.ErrRange
		}
		return uint64(n), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint(), nil
	case reflect.String:
		n, err := strconv.ParseInt(v.String(), 0, 64)
		if err == nil {
			if n < 0 {
				return 0, strconv.ErrRange
			}
			return uint64(n), nil
		}
		if errors.Is(err, strconv.ErrRange) {
			return 0, err
		}
		if f, fErr := strconv.ParseFloat(v.String(), 64); fErr == nil {
			if f < 0 {
				return 0, strconv.ErrRange
			}
			return uint64(f), nil
		}
		return 0, err
	default:
		return 0, strconv.ErrSyntax
	}
}

func ToUniqueIntSlice(a []int) []int {
	m := make(map[int]struct{}, len(a))
	l := make([]int, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		l = append(l, v)
	}
	return l
}

func ToUniqueInt64Slice(a []int64) []int64 {
	m := make(map[int64]struct{}, len(a))
	l := make([]int64, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		l = append(l, v)
	}
	return l
}
