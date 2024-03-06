package response

type FeatureFlag struct {
	Key       string         `json:"key"`
	Status    string         `json:"status"`
	Value     VariationValue `json:"value"`
	ValueType string         `json:"value_type"`
}

type VariationValue struct {
	BooleanValue *bool   `json:"boolean_value,omitempty"`
	StringValue  *string `json:"string_value,omitempty"`
	IsBoolean    *bool   `json:"is_boolean,omitempty"`
	IsString     *bool   `json:"is_string,omitempty"`
}
