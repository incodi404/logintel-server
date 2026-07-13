package utils

type AppError struct {
	Status  int
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func ErrorResponse(status int, code string, message string) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
