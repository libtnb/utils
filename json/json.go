package json

import (
	"bytes"
	"encoding/json"
)

var (
	Encoder = json.NewEncoder
	Decoder = json.NewDecoder
)

// Marshal serializes the given value to a JSON-encoded byte slice.
func Marshal(v any) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := Encoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unmarshal deserializes the given JSON-encoded byte slice into the provided value.
func Unmarshal(data []byte, v any) error {
	dec := Decoder(bytes.NewReader(data))
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

// MarshalString serializes the given value to a JSON-encoded string.
func MarshalString(v any) (string, error) {
	s, err := Marshal(v)
	return string(s), err
}

// UnmarshalString deserializes the given JSON-encoded string into the provided value.
func UnmarshalString(data string, v any) error {
	return Unmarshal([]byte(data), v)
}
