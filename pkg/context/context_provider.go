//go:generate mockgen -source context_provider.go -destination ./mock_context/mock_context_provider/context_provider_gen.go
package context

/**
 * ContextProvider helps set common context values so that client calls do not need to include them in every single call.
 */
type ContextProvider interface {
	/**
	 * Set sets a single condition key using a context.
	 */
	Set(key string, value interface{})

	/**
	 * Unset removes a single condition key from context.
	 */
	Unset(key string)

	/**
	 * Clear removes all existing condition keys from context.
	 */
	Clear()

	/**
	 * Context returns a map of current context.
	 */
	Context() map[string]interface{}
}

func NewContextProvider() ContextProvider {
	return &DefaultContextProvider{}
}
