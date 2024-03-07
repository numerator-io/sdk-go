package request

type FlagListRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type FlagValueByKeyRequest struct {
	Key     string                 `json:"key"`
	Context map[string]interface{} `json:"context"`
}
