package utils

import (
	"net/http"

	"github.com/peterest/go-basic-ecom/types"
)

func HandleSuccessfulAPIResponse(w http.ResponseWriter, data interface{}, message string, status int) {
	if message == "" {
		message = "success"
	}

	if status == 0 {
		status = http.StatusOK
	}

	response := types.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	writeJSON(w, status, response)
}

func HandleFailedAPIResponse(w http.ResponseWriter, status int, message string) {
	if message == "" {
		message = "failed"
	}

	response := types.APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
	}

	writeJSON(w, status, response)
}
