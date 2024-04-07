package config

import (
	_ "embed"

	"github.com/STCraft/DFLoader/config"
)

//go:embed messages.json
var defaultMessages []byte
var messages = config.New("configs", "messages.json", defaultMessages)

// Message returns the formatted message
func Message(message string, args ...any) string {
	return messages.GetMessage(message, args...)
}
