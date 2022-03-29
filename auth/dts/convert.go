package dts

import (
	"encoding/json"
	"io"
	"itsln.com/anonym/errors"
)

func Unmarshal(reader io.Reader, entity interface{}) {
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(entity); err != nil {
		errors.CantDecodeData(err)
	}
}

func Marshal(v interface{}) []byte {
	result, err := json.Marshal(v)
	if err != nil {
		errors.CantEncodeData(err)
	}
	return result
}
