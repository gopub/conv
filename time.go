package conv

import (
	"fmt"
	"strings"
	"time"
)

var dateFormats = []string{
	"2006-1-2", "20060102", "2006/1/2", "2/1/2006",
}

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

func ToDate(i interface{}) (*time.Time, error) {
	return ToDateInLocation(i, time.Local)
}

func ToDateInLocation(i interface{}, loc *time.Location) (*time.Time, error) {
	i = indirect(i)
	s, err := ToString(i)
	if err != nil {
		return nil, fmt.Errorf("cannot convert %#v of type %T to date: %w", i, i, err)
	}
	for _, df := range dateFormats {
		d, err := time.ParseInLocation(df, s, loc)
		if err == nil {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("cannot convert %#v of type %T to date", i, i)
}
