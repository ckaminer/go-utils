package http_helpers

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)

	setResponseHeaders(w)
	w.WriteHeader(status)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	payload := map[string]string{
		"message": message,
	}
	RespondWithJSON(w, status, payload)
}

func RespondWithParsedError(w http.ResponseWriter, err error) {
	if resourceErr, ok := err.(ResourceError); ok {
		RespondWithError(w, resourceErr.Code(), resourceErr.ClientMessage())
	}
	if httpErr, ok := err.(HttpError); ok {
		RespondWithError(w, httpErr.StatusCode, httpErr.Error())
	}
	RespondWithError(w, http.StatusServiceUnavailable, err.Error())
}

func RespondWithStatus(w http.ResponseWriter, status int) {
	setResponseHeaders(w)
	w.WriteHeader(status)
}

func setResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
}
