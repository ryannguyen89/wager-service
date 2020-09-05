package utils

import (
	"encoding/json"
	"log"
)

func JsonSerialize(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("marshal object %v error: %v", v, err)
		return ""
	}

	return string(bytes)
}

func JsonDeserialize(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}
