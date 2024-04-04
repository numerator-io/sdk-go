package response

import (
	"fmt"

	"github.com/c0x12c/numerator-go-sdk/internal/models"
)

/**
 * A data class representing the result of a flag valuation.
 *
 * @param key The unique key of the feature flag.
 * @param value The value of the feature flag.
 * @param reason A map of the reason for the flag valuation. This is an optional field.
 */
type FlagEvaluationDetail[T models.FlagValueType] struct {
	Key    string
	Value  T
	Reason map[string]interface{}
}

/**
 * GetErrorMessage retrieves the error message from the reason map.
 */
func (f *FlagEvaluationDetail[T]) GetErrorMessage() string {
	kind, _ := f.Reason["kind"].(string)
	errorKind, _ := f.Reason["errorKind"].(string)
	if kind != "" && errorKind != "" {
		return kind + ": " + errorKind
	}
	return fmt.Sprintf("%v", f.Reason)
}
