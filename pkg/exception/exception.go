package exception

import "fmt"

type NumeratorException struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewNumeratorException(message string, status int) *NumeratorException {
	return &NumeratorException{Message: message, Status: status}
}

func (e *NumeratorException) Error() string {
	return fmt.Sprintf("Numerator Exception: %s (Status: %d)", e.Message, e.Status)
}
