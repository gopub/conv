package conv

import (
	"fmt"
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
	v = IndirectReadableValue(v)
	switch v.Kind() {
	case reflect.Struct:
		squishStructStringFields(v)
	case reflect.String:
		fmt.Println(v.CanSet(), v.String())
		if v.CanSet() {
			v.SetString(SquishString(v.String()))
		}
	default:
		break
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
