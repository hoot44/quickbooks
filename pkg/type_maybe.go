package quickbooks

import (
	"bytes"
	"encoding/json"
)

type NilString struct {
	Valid  bool
	String string
}

func NilStr(s string) NilString {
	return NilString{true, s}
}

func (e NilString) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return []byte(`null`), nil
	}
	escapedStr, err := json.Marshal(e.String)
	if err != nil {
		return []byte{}, nil
	}
	return []byte(escapedStr), nil
}

func (e *NilString) UnmarshalJSON(d []byte) error {
	if bytes.Compare(d, []byte(`null`)) == 0 {
		(*e).Valid = false
		return nil
	}
	err := json.Unmarshal(d, &e.String)
	if err != nil {
		return err
	}
	e.Valid = true
	return nil
}
