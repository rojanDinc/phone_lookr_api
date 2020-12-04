package controller

type ErrorResponse = map[string]string

func ErrorToMap(err error) ErrorResponse {
	return ErrorResponse{
		"error_message": err.Error(),
	}
}
