package utils

import (
	"fmt"
)

var toasts = GetConfig("toasts.json")

// GetToast parses and returns a toast from a config file.
func GetToast(key string) (title string, content string) {
	object := toasts.Config(key)
	title = object.String("title")
	content = object.String("content")

	return
}

// ToastTitle parses and returns the toast title from the config
func ToastTitle(key string, args ...any) string {
	title, _ := GetToast(key)
	return fmt.Sprintf(title, args...)
}

// ToastTitle parses and returns the toast title from the config
func ToastContent(key string, args ...any) string {
	_, content := GetToast(key)
	return fmt.Sprintf(content, args...)
}
