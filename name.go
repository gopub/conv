package conv

import "strings"

// ToSnake converts s from CamelCase to Snake
func ToSnake(s string) string {
	snake := make([]rune, 0, len(s)+1)
	flag := false
	k := 'a' - 'A'
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if !flag {
				flag = true
				if i > 0 {
					snake = append(snake, '_')
				}
			}
			snake = append(snake, c+k)
		} else {
			flag = false
			snake = append(snake, c)
		}
	}
	return string(snake)
}

// ToCamel converts s from Snake to Camel
func ToCamel(s string) string {
	camel := make([]rune, 0, len(s)+1)
	flag := false
	k := 'A' - 'a'
	for _, c := range s {
		if c == '_' {
			flag = true
			continue
		}

		if flag {
			flag = false
			if c >= 'a' && c <= 'z' {
				camel = append(camel, c+k)
				continue
			}
		}
		camel = append(camel, c)
	}
	return string(camel)
}

// NameChecker check if names can be convertiable
type NameChecker interface {
	CheckName(src, dst string) bool
}

// NameCheckFunc defines func type which implements NameChecker
type NameCheckFunc func(src string, dst string) bool

// CheckName checks if srcName can be converted to dstName
func (f NameCheckFunc) CheckName(srcName, dstName string) bool {
	return f(srcName, dstName)
}

var defaultNameChecker = NameCheckFunc(CheckName)

// CheckName is the default NameChecker
func CheckName(a, b string) bool {
	if a == b {
		return true
	}

	la := strings.ToLower(a)
	lb := strings.ToLower(b)
	switch {
	case la == lb:
		return true
	case strings.ToLower(ToSnake(a)) == lb:
		return true
	case la == strings.ToLower(ToSnake(b)):
		return true
	case strings.ToLower(ToCamel(a)) == lb:
		return true
	case la == strings.ToLower(ToCamel(b)):
		return true
	default:
		return false
	}
}
