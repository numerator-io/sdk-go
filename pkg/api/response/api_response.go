package response

type ApiResponse interface {
	isApiResponse()
}

type SuccessResponse[T interface{}] struct {
	SuccessResponse T `json:"success_response"`
}

func (s SuccessResponse[T]) isApiResponse() {}

type FailureResponse struct {
	Error NumeratorError `json:"error"`
}

func (f FailureResponse) isApiResponse() {}
