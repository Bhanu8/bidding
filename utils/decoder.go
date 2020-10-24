package utils

import (
	"encoding/json"
	"net/http"
)

type ok interface {
	Ok() error
}

func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	if payload, ok := v.(ok); ok {
		return payload.Ok()
	}
	return nil
}

func JustDecode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
