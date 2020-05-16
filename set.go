package conv

import (
	"encoding"
	"fmt"
	"reflect"
)

func SetBytes(target interface{}, b []byte) error {
	target = indirect(target)
	if !reflect.ValueOf(target).CanSet() {
		return fmt.Errorf("target %T is unsetable", target)
	}

	if tu, ok := target.(encoding.TextUnmarshaler); ok {
		err := tu.UnmarshalText(b)
		if err != nil {
			return fmt.Errorf("unmarshal text: %w", err)
		}
		return nil
	}

	if bu, ok := target.(encoding.BinaryUnmarshaler); ok {
		err := bu.UnmarshalBinary(b)
		if err != nil {
			return fmt.Errorf("unmarshal binary: %w", err)
		}
		return nil
	}

	switch v := reflect.ValueOf(target); v.Kind() {
	case reflect.String:
		v.SetString(string(b))
	case reflect.Int64,
		reflect.Int32,
		reflect.Int,
		reflect.Int16,
		reflect.Int8:
		i, err := ToInt64(b)
		if err != nil {
			return fmt.Errorf("parse int: %v", err)
		}
		v.SetInt(i)
	case reflect.Uint64,
		reflect.Uint32,
		reflect.Uint,
		reflect.Uint16,
		reflect.Uint8:
		i, err := ToUint64(b)
		if err != nil {
			return fmt.Errorf("parse uint: %w", err)
		}
		v.SetUint(i)
	case reflect.Float32, reflect.Float64:
		i, err := ToFloat64(b)
		if err != nil {
			return fmt.Errorf("parse float: %w", err)
		}
		v.SetFloat(i)
	case reflect.Bool:
		i, err := ToBool(b)
		if err != nil {
			return fmt.Errorf("parse bool: %w", err)
		}
		v.SetBool(i)
	case reflect.Array:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			v.SetBytes(b)
		} else {
			return fmt.Errorf("cannot set %T", target)
		}
	default:
		return fmt.Errorf("cannot set %T", target)
	}
	return nil
}
