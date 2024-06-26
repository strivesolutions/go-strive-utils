package striveexceptions

import "fmt"

// Exception is a struct that contains the error message and the error code
type Exception struct {
	FullError  error
	StatusCode int         `json:"statusCode"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Details    interface{} `json:"details"`
}

func (e Exception) Error() string {
	if e.FullError != nil {
		return fmt.Sprintf("Code: %d, Message: %s, Details: %s, FullError: %s", e.Code, e.Message, e.Details, e.FullError.Error())
	}

	if e.Details == nil {
		return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
	}

	return fmt.Sprintf("Code: %d, Message: %s, Details: %s", e.Code, e.Message, e.Details)
}
