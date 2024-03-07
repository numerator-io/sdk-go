package response

type ApiResponse interface {
	isApiResponse()
}

type SuccessResponse struct {
	SuccessResponse interface{}
}

func (s SuccessResponse) isApiResponse() {}

type FailureResponse struct {
	Error NumeratorError
}

func (f FailureResponse) isApiResponse() {}
