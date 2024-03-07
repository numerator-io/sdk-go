package response

type PageableResponse interface {
	Count() int64
	Data() interface{}
}
