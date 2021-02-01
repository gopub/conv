package conv

import (
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
)

func ToBytes(i interface{}) ([]byte, error) {
	i = Indirect(i)
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
	return nil, fmt.Errorf("cannot convert %#v of type %T to []byte", i, i)
}

func SetBytes(target interface{}, b []byte) error {
	target = Indirect(target)
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

	if ju, ok := target.(json.Unmarshaler); ok {
		err := ju.UnmarshalJSON(b)
		if err != nil {
			return fmt.Errorf("unmarshal json: %w", err)
		}
		return nil
	}

	v := reflect.ValueOf(target)
	if IsIntValue(v) {
		i, err := ToInt64(b)
		if err != nil {
			return fmt.Errorf("parse int: %v", err)
		}
		v.SetInt(i)
	}

	if IsUintValue(v) {
		i, err := ToUint64(b)
		if err != nil {
			return fmt.Errorf("parse uint: %w", err)
		}
		v.SetUint(i)
	}

	if IsFloatValue(v) {
		i, err := ToFloat64(b)
		if err != nil {
			return fmt.Errorf("parse float: %w", err)
		}
		v.SetFloat(i)
	}

	switch v.Kind() {
	case reflect.String:
		v.SetString(string(b))
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
			return fmt.Errorf("cannot assign %T", target)
		}
	default:
		return fmt.Errorf("cannot assign %T", target)
	}
	return nil
}
