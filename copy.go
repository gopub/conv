package conv

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func GobCopy(dst, src interface{}) error {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(src)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	dec := gob.NewDecoder(&b)
	err = dec.Decode(dst)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}
	return nil
}

func JSONCopy(dst, src interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	err = json.Unmarshal(b, dst)
	if err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	return nil
}

func DeepCopy(dst, src interface{}) error {
	if dst == nil {
		return errors.New("dst cannot be nil")
	}

	if src == nil {
		return errors.New("src cannot be nil")
	}

	dstType := reflect.TypeOf(dst)
	srcType := reflect.TypeOf(src)

	dstKind := IndirectReadableValue(reflect.ValueOf(dst)).Kind()
	srcKind := IndirectReadableValue(reflect.ValueOf(src)).Kind()
	if reflect.PtrTo(srcType).ConvertibleTo(dstType) || (srcKind == dstKind && dstKind == reflect.Struct) {
		err := GobCopy(dst, src)
		if err != nil {
			return fmt.Errorf("json copy: %w", err)
		}
		return nil
	}

	return fmt.Errorf("cannot copy %T to %T", src, dst)
}
