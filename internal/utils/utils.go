package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParsingFromJson(r *http.Request, requestPayload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}
	return json.NewDecoder(r.Body).Decode(requestPayload)
}

func ParsingToJson(w http.ResponseWriter, status int, responseBody any) error {
	w.Header().Add("Content-Type", "application/json")
	if status == http.StatusNoContent {
		w.WriteHeader(status)
		return nil
	}
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(responseBody)
}

func WritingError(w http.ResponseWriter, status int, err error) {
	ParsingToJson(w, status, map[string]string{"Error": err.Error()})
}
