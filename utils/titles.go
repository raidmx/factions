package utils

var titles = GetConfig("titles.json")

// TitleData returns the title data
func TitleData(key string) map[string]any {
	return titles.Get(key).(map[string]any)
}
