package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONRespBody struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
}

// StatusOK for 200
func StatusOK(w http.ResponseWriter, data interface{}) {
	body := &JSONRespBody{
		StatusCode: http.StatusOK,
		Data:       data,
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// StatusCreated for 201
func StatusCreated(w http.ResponseWriter) {
	body := &JSONRespBody{
		StatusCode: http.StatusCreated,
		Data:       "new item created",
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

// StatusCreated for 204
func StatusNoContent(w http.ResponseWriter) {
	body := &JSONRespBody{
		StatusCode: http.StatusNoContent,
		Data:       "row updated",
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
	w.Write(b)
}

// BadRequest for 400
func BadRequest(w http.ResponseWriter) {
	body := &JSONRespBody{
		StatusCode: http.StatusBadRequest,
		Data:       "Invalid payload body",
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(b)
}

// NotFound for 404
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

// UnprocessableEntity for 422
func UnprocessableEntity(w http.ResponseWriter) {
	body := &JSONRespBody{
		StatusCode: http.StatusUnprocessableEntity,
		Data:       "Invalid payload body",
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(b)
}

// InternalServerError for 500
func InternalServerError(w http.ResponseWriter) {
	body := &JSONRespBody{
		StatusCode: http.StatusInternalServerError,
		Data:       "Internal Server Error",
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}
