package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type jsonResponse struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

type jsonErrorResponse struct {
	Status interface{} `json:"status"`
	Errors []string    `json:"errors"`
}

type newJsonErrorResponse struct {
	Status interface{}   `json:"status"`
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type failedCallbackResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type successCallbackResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BadRequestCallback struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, code int, errors []error) {
	var errStr []string
	for _, err := range errors {
		space := regexp.MustCompile(`\s+`)
		errReplace := space.ReplaceAllString(err.Error(), " ")
		errStr = append(errStr, errReplace)
	}

	jsonResponse := jsonErrorResponse{
		Status: Status{
			Code:    code,
			Message: "Error",
		},
		Errors: errStr,
	}
	response, _ := json.Marshal(jsonResponse)

	w.Header().Set("Content-Type", "usecase/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func RespondWithErrorV2(w http.ResponseWriter, code int, errors []ErrorDetail) {

	jsonResponse := newJsonErrorResponse{
		Status: Status{
			Code:    code,
			Message: "Error",
		},
		Errors: errors,
	}
	response, _ := json.Marshal(jsonResponse)

	w.Header().Set("Content-Type", "usecase/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(jsonResponse{
		Status: Status{
			Code:    code,
			Message: "Success",
		},
		Data: payload,
		Meta: nil,
	})

	w.Header().Set("Content-Type", "usecase/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func RespondWithJSONV2(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "usecase/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func ResponseFailedCallback(w http.ResponseWriter, code int, payload BadRequestCallback) {
	response, _ := json.Marshal(failedCallbackResponse{
		Code:    payload.Code,
		Message: payload.Message,
	})

	w.Header().Set("Content-Type", "usecase/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func RespondSuccessCallbackWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(successCallbackResponse{
		Code:    "200",
		Message: "Success",
		Data:    payload,
	})

	w.Header().Set("Content-Type", "usecase/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func NewResponseError(err error) []byte {
	jsonResponse := jsonErrorResponse{
		Status: Status{
			Code:    0,
			Message: err.Error(),
		},
	}
	response, _ := json.Marshal(jsonResponse)
	return response
}

func RespondWithSingleError(w http.ResponseWriter, statusCode int, err error) {
	response := NewResponseError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(response)
	if err != nil {
		return
	}
}
