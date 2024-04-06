package utils

import "fmt"

var messages = GetConfig("messages.json")

// Message returns the formatted message
func Message(message string, args ...any) string {
	msg := messages.String(message)
	return fmt.Sprintf(msg, args...)
}
