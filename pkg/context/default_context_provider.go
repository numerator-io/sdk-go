package context

import (
	"sync"
)

type DefaultContextProvider struct {
	values sync.Map
}

func (c *DefaultContextProvider) Set(key string, value interface{}) {
	c.values.Store(key, value)
}

func (c *DefaultContextProvider) Unset(key string) {
	c.values.Delete(key)
}

func (c *DefaultContextProvider) Clear() {
	c.values.Range(func(key, _ interface{}) bool {
		c.values.Delete(key)
		return true
	})
}

func (c *DefaultContextProvider) Context() map[string]interface{} {
	result := make(map[string]interface{})
	c.values.Range(func(key, value interface{}) bool {
		result[key.(string)] = value
		return true
	})
	return result
}
