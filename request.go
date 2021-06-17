package common

import (
	"encoding/json"
	"io"

	"github.com/mailru/easyjson"
)

func Bind(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		exc := NewSingleMessageException(
			EINVALID,
			"",
			"Invalid Payload",
			err,
		)
		return exc
	}
	return nil
}

func BindEasyJSON(r io.Reader, v easyjson.Unmarshaler) error {
	err := easyjson.UnmarshalFromReader(r, v)
	if err != nil {
		exc := NewSingleMessageException(
			EINVALID,
			"",
			"Invalid Payload",
			err,
		)
		return exc
	}
	return nil
}
