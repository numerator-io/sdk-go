package request

type FlagByKeyRequest struct {
	Key     string
	Context map[string]interface{}
}

type FlagValueByKeyRequest struct {
	Key     string
	Context map[string]interface{}
}
