package conv

import (
	"encoding/json"
	"log"
)

func MustJSONString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return string(b)
}

func MustJSONBytes(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return b
}
