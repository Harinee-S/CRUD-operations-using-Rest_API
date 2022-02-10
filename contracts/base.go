package contracts

import (
	"encoding/json"
	"net/http"
)

type errorList struct {
	Message string `json:"message,omitempty"`
}

type baseResponse struct {
	Success bool        `json:"success"`
	Errors  []errorList `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(rw http.ResponseWriter, responseData interface{}, success string) {
	successDetails := successObjects[success]

	rw.WriteHeader(successDetails.status)
	rw.Header().Set("Content-Type", "application/json")
	response := baseResponse{
		Success: true,
		Data:    responseData,
	}

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)

	return
}

func ErrorResponse(rw http.ResponseWriter, errors []string, status string) {
	error := errorObjects[status]
	rw.WriteHeader(error.status)
	rw.Header().Set("Content-Type", "application/json")

	var errorSlice []errorList

	for _, error := range errors {
		errorSlice = append(errorSlice, errorList{Message: error})
	}

	response := baseResponse{
		Success: false,
		Errors:  errorSlice,
	}

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)
	return
}
