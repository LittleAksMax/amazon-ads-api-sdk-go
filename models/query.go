package models

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strings"
)

// toQueryValues converts any struct with 'query' tags into url.Values
// It uses reflection to iterate over struct fields and extracts their values
// based on the 'query' struct tag. Supports string and []string types only.
func toQueryValues(v interface{}) url.Values {
	queryParameters := url.Values{}

	val := reflect.ValueOf(v)
	// If a pointer is passed, dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return queryParameters
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		queryTag := field.Tag.Get("query")

		// Skip fields without a query tag
		if queryTag == "" {
			continue
		}

		fieldVal := val.Field(i)

		// Handle string type
		if fieldVal.Kind() == reflect.String {
			if str := fieldVal.String(); str != "" {
				queryParameters.Add(queryTag, str)
			}
		}

		// Handle []string type
		if fieldVal.Kind() == reflect.Slice && fieldVal.Type().Elem().Kind() == reflect.String {
			if fieldVal.Len() > 0 {
				var strSlice []string
				for j := 0; j < fieldVal.Len(); j++ {
					strSlice = append(strSlice, fieldVal.Index(j).String())
				}
				queryParameters.Add(queryTag, strings.Join(strSlice, ","))
			}
		}
	}

	return queryParameters
}

// toJSONBodyOptions converts any struct with 'json' tags into a JSON body format
// For array fields, it wraps them in an "includes" object with the json tag as key
// For nested objects, it uses Marshal directly
func toJSONBodyOptions(v interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(v)
	// If a pointer is passed, dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return result
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")

		// Skip fields without a json tag
		if jsonTag == "" {
			continue
		}

		// Parse the json tag to extract just the key name (before any comma)
		tagKey := strings.Split(jsonTag, ",")[0]
		if tagKey == "" {
			continue
		}

		fieldVal := val.Field(i)

		// Skip zero/empty values
		if fieldVal.IsZero() {
			continue
		}

		// Handle slice/array types
		if fieldVal.Kind() == reflect.Slice {
			// Create the includes wrapper object
			includesObj := make(map[string]interface{})

			// For slices, we just add the slice directly to includes
			var sliceData []interface{}
			for j := 0; j < fieldVal.Len(); j++ {
				sliceData = append(sliceData, fieldVal.Index(j).Interface())
			}
			includesObj["includes"] = sliceData

			result[tagKey] = includesObj
			continue
		}

		// Handle nested objects/structs
		if fieldVal.Kind() == reflect.Struct {
			// Use json.Marshal for nested objects
			data, err := json.Marshal(fieldVal.Interface())
			if err == nil {
				var obj interface{}
				_ = json.Unmarshal(data, &obj)
				result[tagKey] = obj
			}
			continue
		}

		// Handle primitive types
		result[tagKey] = fieldVal.Interface()
	}

	return result
}
