package utils

var uis = GetConfig("uis.json")

// GetUI returns the UI data
func GetUI(ui string) map[string]any {
	return uis.Get(ui).(map[string]any)
}
