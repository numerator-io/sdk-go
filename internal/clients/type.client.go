package clients

import (
	"time"
)

// ConfigClient represents configuration client.
type ConfigClient struct {
	APIKey  string
	BaseURL string
}

// ErrorResponse represents error response structure.
type ErrorResponse struct {
	Message     string
	ErrorCode   string
	ErrorStatus int
}

// ApiResponse represents API response structure.
type ApiResponse[T any] struct {
	Data  *T
	Error *ErrorResponse
}

// ApiRequestOptions represents API request options.
type ApiRequestOptions struct {
	Endpoint string
}

// ApiClientInterface represents API client interface.
type ApiClientInterface interface {
	Request[T any](apiRequestOptions ApiRequestOptions) (*ApiResponse[T], error)
}

// PaginationRequest represents pagination request structure.
type PaginationRequest struct {
	Page int
	Size int
}

// PaginationResponse represents pagination response structure.
type PaginationResponse[T any] struct {
	Count int
	Data  []T
}

// FeatureFlagConfigListingRequest represents feature flag config listing request structure.
type FeatureFlagConfigListingRequest struct {
	PaginationRequest
}

// FeatureFlagValueByKeyRequest represents feature flag value by key request structure.
type FeatureFlagValueByKeyRequest struct {
	Key     string
	Context map[string]interface{}
}

// FeatureFlagConfigListingResponse represents feature flag config listing response structure.
type FeatureFlagConfigListingResponse struct {
	PaginationResponse[FeatureFlagConfig]
}

// VariationKeyType represents variation key type.
type VariationKeyType string

// FlagStatusEnum represents flag status enumeration.
type FlagStatusEnum string

const (
	// ON status.
	ON FlagStatusEnum = "ON"
	// OFF status.
	OFF FlagStatusEnum = "OFF"
)

// FlagValueTypeEnum represents flag value type enumeration.
type FlagValueTypeEnum string

const (
	// BOOLEAN value type.
	BOOLEAN FlagValueTypeEnum = "BOOLEAN"
	// STRING value type.
	STRING FlagValueTypeEnum = "STRING"
	// LONG value type.
	LONG FlagValueTypeEnum = "LONG"
	// DOUBLE value type.
	DOUBLE FlagValueTypeEnum = "DOUBLE"
)

// FeatureFlagConfig represents feature flag config structure.
type FeatureFlagConfig struct {
	ID                  string
	Name                string
	Key                 string
	OrganizationID      string
	ProjectID           string
	Status              FlagStatusEnum
	Description         *string
	DefaultOnVariationID string
	DefaultOffVariationID string
	ValueType           FlagValueTypeEnum
	CreatedAt           time.Time
}

type FeatureFlagValue[T any] struct {
	Key       string
	Status    FlagStatusEnum
	Value     T
	ValueType FlagValueTypeEnum
}

type VariationValue struct {
	StringValue  *string
	BooleanValue *bool
	LongValue    *int
	DoubleValue  *float64
}
