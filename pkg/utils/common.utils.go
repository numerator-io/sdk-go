package utils

import (
	"encoding/json"
	"errors"
	"time"
)

// Deep copy an object using JSON.
func deepCopy(obj interface{}) interface{} {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	var copy interface{}
	err = json.Unmarshal(data, &copy)
	if err != nil {
		return nil
	}
	return copy
}

// Asynchronous sleep function using Promises.
func sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

// Map an array of objects to a map using a specific key.
func mapArrayToMap(array []map[string]interface{}) map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})
	for _, item := range array {
		key := item["key"].(string)
		result[key] = item
	}
	return result
}

// Create a promise with a timeout.
func withTimeout(promise chan interface{}, timeout int) (interface{}, error) {
	select {
	case result := <-promise:
		return result, nil
	case <-time.After(time.Duration(timeout) * time.Millisecond):
		return nil, errors.New("operation timed out")
	}
}

// Convert snakecase map to camelcase map.
func snakeToCamel(obj map[string]interface{}) map[string]interface{} {
	camelObj := make(map[string]interface{})
	for key, value := range obj {
		camelKey := toCamelCase(key)
		switch v := value.(type) {
		case map[string]interface{}:
			camelObj[camelKey] = snakeToCamel(v)
		default:
			camelObj[camelKey] = v
		}
	}
	return camelObj
}

// Helper function to convert snake case to camel case.
func toCamelCase(s string) string {
	result := make([]rune, 0, len(s))
	upper := false
	for _, r := range s {
		if r == '_' {
			upper = true
			continue
		}
		if upper {
			result = append(result, []rune{r - 32}...)
			upper = false
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
