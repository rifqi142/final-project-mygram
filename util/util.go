package util

import "final-project-mygram/model"

func CreateResponse(isSuccess bool, data any, errorMessage string, message string) model.Response {
	return model.Response{
		Success: isSuccess,
		Data:    data,
		Error:   errorMessage,
		Message: message,
	}
}