package dts

import (
	"encoding/json"
	"github.com/smbody/anonym/errors"
	"io"
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
