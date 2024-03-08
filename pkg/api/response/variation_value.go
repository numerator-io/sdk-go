package response

type VariationValue struct {
	BooleanValue bool    `json:"boolean_value,omitempty"`
	StringValue  string  `json:"string_value,omitempty"`
	LongValue    int64   `json:"long_value,omitempty"`
	DoubleValue  float64 `json:"double_value,omitempty"`
}
