package common

import (
	"encoding/json"
	"net/http"

	"github.com/mailru/easyjson"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func NewResponse(status int, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (r *Response) SendEasyJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	b, err := easyjson.Marshal(r)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func (r *Response) SendJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	return json.NewEncoder(w).Encode(r)
}

func SendError(w http.ResponseWriter, err error) {
	if exc, ok := err.(*Exception); ok {
		code := ExceptionCode(err)
		errorResponse := &ErrorResponse{
			Status:  ExceptionCodeToHTTPStatusCode(code),
			Message: code,
			Errors:  ExceptionMessage(err),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errorResponse.Status)

		json.NewEncoder(w).Encode(errorResponse)

		if errorResponse.Status >= 500 {
			if exc.SourceLocation == nil {
				exc.SourceLocation = new(SourceLocation)
			}
		}
	} else {
		errorResponse := &ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: EINTERNAL,
			Errors:  []string{"Server Error, Try Again later."},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errorResponse.Status)

		json.NewEncoder(w).Encode(errorResponse)
	}
}

func NewErrorResponse(err error) *ErrorResponse {
	if e, ok := err.(*Exception); ok {
		return &ErrorResponse{
			Status:  ExceptionCodeToHTTPStatusCode(e.Code),
			Message: e.Code,
			Errors:  e.Message,
		}
	}

	return &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: EINTERNAL,
		Errors:  []string{"Server Error, Try Again later."},
	}
}

func (r *ErrorResponse) SendJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	return json.NewEncoder(w).Encode(r)
}
