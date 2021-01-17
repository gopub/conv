package conv

import (
	"reflect"
	"regexp"
	"strings"
)

var whitespaceRegexp = regexp.MustCompile(`[ \t\n\r]+`)

// SquishString returns the string
// first removing all whitespace on both ends of the string,
// and then changing remaining consecutive whitespace groups into one space each.
func SquishString(s string) string {
	s = strings.TrimSpace(s)
	s = whitespaceRegexp.ReplaceAllString(s, " ")
	return s
}

func SquishStringFields(i interface{}) {
	squishStringFields(reflect.ValueOf(i))
}

func squishStringFields(v reflect.Value) {
	v = indirectSrcVal(v)
	switch v.Kind() {
	case reflect.Struct:
		squishStructStringFields(v)
	case reflect.Map:
		squishMapStringValues(v)
	default:
		break
	}
}

func squishMapStringValues(v reflect.Value) {
	for _, k := range v.MapKeys() {
		val := v.MapIndex(k)
		switch val.Kind() {
		case reflect.String:
			val.SetString(SquishString(val.String()))
		case reflect.Struct:
			squishStructStringFields(val)
		case reflect.Map:
			squishMapStringValues(val)
		case reflect.Ptr:
			SquishStringFields(val.Elem().Interface())
		default:
			break
		}
	}
}

func squishStructStringFields(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		if !fv.IsValid() || !fv.CanSet() {
			continue
		}
		switch fv.Kind() {
		case reflect.String:
			fv.SetString(SquishString(fv.String()))
		case reflect.Struct:
			squishStructStringFields(fv)
		case reflect.Map:
			squishMapStringValues(fv)
		case reflect.Ptr, reflect.Interface:
			if fv.IsNil() {
				break
			}
			squishStringFields(fv.Elem())
		default:
			break
		}
	}
}
