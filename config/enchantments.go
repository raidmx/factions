package config

import (
	_ "embed"

	"github.com/STCraft/DFLoader/config"
)

//go:embed enchantments.json
var defaultEnchantments []byte
var enchantments = config.New("./configs/enchantments.json", defaultEnchantments)

// EnchantmentName returns the enchantment name
func EnchantmentName(key string) string {
	data := enchantments.GetObject(key)
	return data.GetString("name")
}

// EnchantmentDescription returns the enchantment description
func EnchantmentDescription(key string) string {
	data := enchantments.GetObject(key)
	return data.GetString("description")
}

// MaxEnchantmentLevel returns the maximum enchantment level
func MaxEnchantmentLevel(key string) int {
	data := enchantments.GetObject(key)
	return data.GetInt("max_level")
}

// MinEnchantmentCost returns the minimum enchantment cost
func MinEnchantmentCost(key string) int {
	data := enchantments.GetObject(key)
	return data.GetInt("min_cost")
}

// EnchantmentCost returns the maximum enchantment cost
func MaxEnchantmentCost(key string) int {
	data := enchantments.GetObject(key)
	return data.GetInt("max_cost")
}

// EnchantmentProperty returns the custom enchantment property value
func EnchantmentProperty[T any](key string, property string) T {
	data := enchantments.GetObject(key)
	return data.Get(property).(T)
}
