package response

type NumeratorError struct {
	Message    string `json:"message,omitempty"`
	HttpStatus int    `json:"http_status,omitempty"`
}

func (e NumeratorError) Error() string {
	return e.Message
}
