package utils

import (
	"encoding/json"
	"sync"
)

// Substring returns a part of the string specified
func Substring(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

// LenSyncMap ...
func LenSyncMap(m *sync.Map) int {
	var i int
	m.Range(func(k, v interface{}) bool {
		i++
		return true
	})
	return i
}

// Encode encodes the provided value into JSON and returns it
func Encode(val any) string {
	json, err := json.MarshalIndent(val, "", "  ")

	if err != nil {
		panic(err)
	}

	return string(json)
}
