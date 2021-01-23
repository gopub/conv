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

func VarargsToURLValues(keyAndValues ...interface{}) (url.Values, error) {
	uv := url.Values{}
	keys, vals, err := VarargsToSlice(keyAndValues)
	if err != nil {
		return nil, err
	}
	for i, k := range keys {
		vs, err := ToString(vals[i])
		if err != nil {
			return nil, err
		}
		if vs != "" {
			uv.Add(k, vs)
		}
	}
	return uv, nil
}

func VarargsToSlice(keyValues ...interface{}) (keys []string, values []interface{}, err error) {
	n := len(keyValues)
	if n%2 != 0 {
		err = errors.New("keyValues should be pairs of (string, interface{})")
		return
	}

	keys, values = make([]string, 0, n/2), make([]interface{}, 0, n/2)
	for i := 0; i < n/2; i++ {
		if k, ok := keyValues[2*i].(string); !ok {
			err = fmt.Errorf("keyValues[%d] isn't convertible to string", i)
			return
		} else if keyValues[2*i+1] == nil {
			err = fmt.Errorf("keyValues[%d] is nil", 2*i+1)
			return
		} else {
			keys = append(keys, k)
			values = append(values, keyValues[2*i+1])
		}
	}
	return
}
