package config

import (
	_ "embed"

	"github.com/STCraft/DFLoader/config"
)

//go:embed titles.json
var defaultTitles []byte
var titles = config.New("configs", "titles.json", defaultTitles)

// TitleData returns the title data
func TitleData(key string) map[string]any {
	return titles.Get(key).(map[string]any)
}
