package utils

import (
	"encoding/json"
	"log"
)

// ToJSON converts any value to a JSON string
func ToJSON(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error marshaling to JSON: %v", err)
		return ""
	}
	return string(bytes)
}

// FromJSON converts a JSON string to a struct
func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
