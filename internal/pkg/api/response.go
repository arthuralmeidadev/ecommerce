package api

import (
	"encoding/json"
	"net/http"
)

type response struct {
	w http.ResponseWriter
}

func (res *response) WriteString(s string) {
	res.w.WriteHeader(200)
	res.w.Write([]byte(s))
}

func (res *response) WriteJSON(v any) {
	res.w.WriteHeader(200)
	res.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res.w).Encode(v)
}

func (res *response) WriteStatusMessage(status int, msg string) {
	res.w.WriteHeader(status)
	res.w.Header().Set("Content-Type", "application/json")
	data := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: msg,
	}
	json.NewEncoder(res.w).Encode(&data)
	return
}

func (res *response) BadRequestError(msg string) {
	http.Error(res.w, msg, http.StatusBadRequest)
	return
}

func (res *response) NotFoundError(msg string) {
	http.Error(res.w, msg, http.StatusNotFound)
	return
}
