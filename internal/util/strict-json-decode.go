package util

import (
	"bytes"
	"encoding/json"
)

func StrictJSONDecode(raw []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}
