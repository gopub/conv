package conv

import (
	"encoding"
	"errors"
	"fmt"
	"log"
	"reflect"
)

// Assign fill src's underlying value and fields with dst
func Assign(dst interface{}, src interface{}) error {
	return AssignC(dst, src, defaultNameChecker)
}

// AssignC assigns value with name checker
func AssignC(dst interface{}, src interface{}, checker NameChecker) error {
	if dst == nil || src == nil || checker == nil {
		panic(fmt.Sprintf("Cannot accept nil arguments: %v, %v, %v", dst, src, checker))
	}
	_ = JSONCopy(dst, src)
	dv := IndirectWritableValue(reflect.ValueOf(dst), false)
	// dv must be a nil pointer or a valid value
	err := assign(dv, reflect.ValueOf(src), checker)
	if err != nil {
		return fmt.Errorf("cannot assign %T to %T: %w", src, dv.Interface(), err)
	}
	if err = Validate(dst); err != nil {
		return fmt.Errorf("cannot validate: %w", err)
	}
	return nil
}

// dst is valid value or pointer to value
func assign(dst reflect.Value, src reflect.Value, nm NameChecker) error {
	src = IndirectReadableValue(src)
	dv := IndirectWritableValue(dst, true)
	switch dv.Kind() {
	case reflect.Bool:
		b, err := ToBool(src.Interface())
		if err != nil {
			return fmt.Errorf("parse bool: %w", err)
		}
		dv.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := ToInt64(src.Interface())
		if err != nil {
			return fmt.Errorf("parse int64: %w", err)
		}
		dv.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := ToUint64(src.Interface())
		if err != nil {
			return fmt.Errorf("parse uint64: %w", err)
		}
		dv.SetUint(i)
	case reflect.Float32, reflect.Float64:
		i, err := ToFloat64(src.Interface())
		if err != nil {
			return fmt.Errorf("parse float64: %w", err)
		}
		dv.SetFloat(i)
	case reflect.String:
		if src.Kind() != reflect.String {
			return errors.New("source value is not string")
		}
		dv.SetString(src.String())
	case reflect.Slice:
		if src.Kind() != reflect.Slice {
			return errors.New("source value is not slice")
		}
		l := reflect.MakeSlice(dv.Type(), src.Len(), src.Cap())
		for i := 0; i < src.Len(); i++ {
			err := assign(l.Index(i), src.Index(i), nm)
			if err != nil {
				return fmt.Errorf("cannot assign field at index %d: %w", i, err)
			}
		}
		dv.Set(l)
	case reflect.Map:
		if k := src.Kind(); k != reflect.Map {
			return fmt.Errorf("cannot assign %v to map", k)
		}
		if err := mapToMap(dv, src, nm); err != nil {
			return fmt.Errorf("cannot assign map to map: %w", err)
		}
	case reflect.Struct:
		err := valueToStruct(dv, src, nm)
		if err != nil {
			return fmt.Errorf("assign to struct: %w", err)
		}
	case reflect.Interface:
		// if i is a pointer to an interface, then ValueOf(i).Elem().Kind() is reflect.Interface
		switch ev := dv.Elem(); ev.Kind() {
		case reflect.Map:
			if k := src.Kind(); k != reflect.Map {
				return fmt.Errorf("cannot assign %v to map", k)
			}
			pv := reflect.New(ev.Type())
			err := mapToMap(pv.Elem(), src, nm)
			if err != nil {
				return fmt.Errorf("cannot assign map to interface(map): %w", err)
			}
			dv.Set(pv.Elem())
		case reflect.Struct:
			pv := reflect.New(ev.Type())
			if err := valueToStruct(pv.Elem(), src, nm); err != nil {
				return fmt.Errorf("cannot assign to interface(struct): %w", err)
			}
			dv.Set(pv.Elem())
		default:
			panic(fmt.Sprintf("unknown interface(kind)=%s", ev))
		}
	default:
		panic(fmt.Sprintf("unknown kind=%v", dv.Kind()))
	}

	if dst.Kind() == reflect.Ptr && dst.IsNil() {
		dst.Set(dv.Addr())
	}
	return nil
}

func valueToStruct(dst reflect.Value, src reflect.Value, nm NameChecker) error {
	switch k := src.Kind(); k {
	case reflect.Map:
		return wrapError(mapToStruct(dst, src, nm), "mapToStruct")
	case reflect.Struct:
		return wrapError(structToStruct(dst, src, nm), "structToStruct")
	case reflect.String:
		if dst.CanInterface() {
			if u, ok := dst.Interface().(encoding.TextUnmarshaler); ok && u != nil {
				err := u.UnmarshalText([]byte(src.String()))
				if err != nil {
					return fmt.Errorf("cannot unmarshal text into %v: %w", dst.Type(), err)
				}
				return nil
			}

			if dst.CanAddr() && dst.Addr().CanInterface() {
				if u, ok := dst.Addr().Interface().(encoding.TextUnmarshaler); ok && u != nil {
					err := u.UnmarshalText([]byte(src.String()))
					if err != nil {
						return fmt.Errorf("cannot unmarshal text into pointer to %v: %w", dst.Type(), err)
					}
					return nil
				}
			}
		}
		return fmt.Errorf("src is %v instead of struct or map", k)
	case reflect.Ptr:
		if src.IsNil() {
			return nil
		}
		return valueToStruct(dst, src.Elem(), nm)
	default:
		return fmt.Errorf("src is %v instead of struct or map", k)
	}
	return nil
}

func mapToMap(dst reflect.Value, src reflect.Value, nm NameChecker) error {
	if !src.Type().Key().AssignableTo(dst.Type().Key()) {
		return fmt.Errorf("cannot assign %s to %s", src.Type().Key(), dst.Type().Key())
	}

	if dst.IsNil() {
		dst.Set(reflect.MakeMap(dst.Type()))
	}

	de := dst.Type().Elem()
	canAssign := src.Type().Elem().AssignableTo(de)
	for _, k := range src.MapKeys() {
		switch {
		case canAssign:
			dst.SetMapIndex(k, src.MapIndex(k))
		case de.Kind() == reflect.Ptr:
			kv := reflect.New(de.Elem())
			err := assign(kv, src.MapIndex(k), nm)
			if err != nil {
				log.Printf("Cannot assign: %v", k.Interface())
				break
			}
			dst.SetMapIndex(k, kv)
		default:
			kv := reflect.New(de)
			err := assign(kv, src.MapIndex(k), nm)
			if err != nil {
				log.Printf("Cannot assign: %v", k.Interface())
				break
			}
			dst.SetMapIndex(k, kv.Elem())
		}
	}
	return nil
}

func mapToStruct(dst reflect.Value, src reflect.Value, nm NameChecker) error {
	if k := src.Type().Key().Kind(); k != reflect.String {
		return fmt.Errorf("src key is %s intead of string", k)
	}

	for i := 0; i < dst.NumField(); i++ {
		fv := dst.Field(i)
		if fv.IsValid() == false || fv.CanSet() == false {
			continue
		}

		ft := dst.Type().Field(i)
		if ft.Anonymous {
			err := assign(fv, src, nm)
			if err != nil {
				log.Printf("Cannot assign %s: %v", ft.Name, err)
			}
			continue
		}

		for _, key := range src.MapKeys() {
			if !nm.CheckName(key.String(), ft.Name) {
				continue
			}

			fsv := src.MapIndex(key)
			if !fsv.IsValid() {
				log.Printf("Invalid value for %s", ft.Name)
				continue
			}

			if fsv.Interface() == nil {
				continue
			}

			err := assign(fv, reflect.ValueOf(fsv.Interface()), nm)
			if err != nil {
				log.Printf("Cannot assign %s: %v", ft.Name, err)
			}
			break
		}
	}
	return nil
}

func structToStruct(dst reflect.Value, src reflect.Value, nm NameChecker) error {
	for i := 0; i < dst.NumField(); i++ {
		fv := dst.Field(i)
		if fv.IsValid() == false || fv.CanSet() == false {
			continue
		}

		ft := dst.Type().Field(i)
		if ft.Anonymous {
			if err := assign(fv, src, nm); err != nil {
				log.Printf("Cannot assign anonymous %s: %v", ft.Name, err)
			}
			continue
		}

		for i := 0; i < src.NumField(); i++ {
			sfv := src.Field(i)
			sfName := src.Type().Field(i).Name
			if !sfv.IsValid() || sfv.Interface() == nil {
				continue
			}

			if !isExported(sfName) || !nm.CheckName(sfName, ft.Name) {
				continue
			}

			err := assign(fv, reflect.ValueOf(sfv.Interface()), nm)
			if err != nil {
				log.Printf("Cannot assign %s to %s: %v", ft.Name, sfName, err)
			}
			break
		}
	}

	for i := 0; i < src.NumField(); i++ {
		sfv := src.Field(i)
		sfName := src.Type().Field(i).Name
		if !sfv.IsValid() || (sfv.CanInterface() && sfv.Interface() == nil) || sfv.IsZero() || !isExported(sfName) {
			continue
		}

		if src.Type().Field(i).Anonymous {
			_ = assign(dst, reflect.ValueOf(sfv.Interface()), nm)
		}
	}
	return nil
}

func isExported(name string) bool {
	return name != "" && name[0] >= 'A' && name[0] <= 'Z'
}
