package conv

import (
	"fmt"
	"strings"
	"time"
)

// ToDuration converts an interface to a time.Duration type.
func ToDuration(i interface{}) (time.Duration, error) {
	i = indirect(i)

	if s, err := ToString(i); err == nil {
		if strings.ContainsAny(s, "nsuµmh") {
			return time.ParseDuration(s)
		} else {
			return time.ParseDuration(s + "ns")
		}
	}

	if n, err := ToInt64(i); err == nil {
		return time.Duration(n), nil
	}

	return 0, fmt.Errorf("cannot convert %#v of type %T to duration", i, i)
}
