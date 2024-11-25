package utils

import "encoding/json"

func GenerateErrorResponse(message string, errorDetails string) string {
	errorResponse := map[string]interface{}{
		"error":         true,
		"message":       message,
		"error_details": errorDetails,
	}
	response, _ := json.Marshal(errorResponse)
	return string(response)
}
