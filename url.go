package conv

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"
)

func ToURLValues(i interface{}) (url.Values, error) {
	i = IndirectToStringerOrError(i)
	if i == nil {
		return nil, errNilValue
	}
	switch v := i.(type) {
	case url.Values:
		return v, nil
	}

	b, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("cannot convert %#v of type %T to url.Values", i, i)
	}
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, fmt.Errorf("cannot convert %#v of type %T to url.Values", i, i)
	}
	uv := url.Values{}
	for k, v := range m {
		uv.Set(k, fmt.Sprint(v))
	}
	return uv, nil
}

func MustURLValues(i interface{}) url.Values {
	v, err := ToURLValues(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func JoinURLPath(a ...string) string {
	if len(a) == 0 {
		return ""
	}
	p := path.Join(a...)
	p = strings.Replace(p, ":/", "://", 1)
	i := strings.Index(p, "://")
	s := p
	if i >= 0 {
		s = p[i:]
		l := strings.Split(s, "/")
		for i, v := range l {
			l[i] = url.PathEscape(v)
		}
		p = p[:i] + path.Join(l...)
	}
	return p
}

func ToURL(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("parse: %w", err)
	}
	if u.Scheme == "" {
		return "", errors.New("missing schema")
	}
	if u.Host == "" {
		return "", errors.New("missing host")
	}
	return u.String(), nil
}

func IsURL(s string) bool {
	u, _ := ToURL(s)
	return u != ""
}
