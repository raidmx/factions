package config

import (
	_ "embed"
	"fmt"

	"github.com/STCraft/DFLoader/config"
)

//go:embed toasts.json
var defaultToasts []byte
var toasts = config.New("configs", "toasts.json", defaultToasts)

// GetToast parses and returns a toast from a config file.
func GetToast(key string) (title string, content string) {
	return toasts.GetToast(key)
}

func ToastTitle(key string) string {
	toast := toasts.GetObject(key)
	return toast.GetString("title")
}

func ToastContent(key string, args ...any) string {
	toast := toasts.GetObject(key)
	return fmt.Sprintf(toast.GetString("content"), args...)
}
