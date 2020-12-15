package conv

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
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
