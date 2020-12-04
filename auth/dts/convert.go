package dts

import (
	"encoding/json"
	"github.com/smbody/anonym/errors"
	"io"
)

func Unmarshal(reader io.Reader, entity interface{}) {
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(entity); err != nil {
		errors.Throw(errors.CantDecodeData)
	}
}

func Marshal(v interface{}) []byte {
	if result, err := json.Marshal(v); err == nil {
		return result
	}
	errors.Throw(errors.CantEncodeData)
	return nil
}


