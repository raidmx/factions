package config

import (
	_ "embed"

	"github.com/STCraft/DFLoader/config"
)

//go:embed uis.json
var defaultUIs []byte
var uis = config.New("configs", "uis.json", defaultUIs)

// GetUI returns the UI data
func GetUI(ui string) map[string]any {
	return uis.Get(ui).(map[string]any)
}
