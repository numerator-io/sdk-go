package request

type FlagListRequest struct {
	Page int
	Size int
}

type FlagValueByKeyRequest struct {
	Key     string
	Context map[string]interface{}
}
