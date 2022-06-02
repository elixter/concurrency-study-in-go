package utils

import (
	"encoding/json"
	"net/http"
)

type ErrMsg string
type AdditionalData interface {
	apply(*map[string]interface{})
}

func WriteJson(w http.ResponseWriter, code int, data interface{}) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(json)

	return err
}

func WriteError(w http.ResponseWriter, code int, errMsg ErrMsg, additionalColumns ...AdditionalData) error {
	body := make(map[string]interface{})
	body["statusCode"] = code
	body["errMsg"] = errMsg

	for _, c := range additionalColumns {
		c.apply(&body)
	}

	return WriteJson(w, code, body)
}