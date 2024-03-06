package exception

type NumeratorException struct {
	message string
}

func NewNumeratorException(message string) *NumeratorException {
	return &NumeratorException{message: message}
}

func (e *NumeratorException) Error() string {
	return e.message
}
