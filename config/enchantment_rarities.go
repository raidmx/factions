package config

import (
	_ "embed"

	"github.com/STCraft/DFLoader/config"
)

//go:embed enchantment_rarities.json
var defaultEnchantmentRarities []byte
var enchantmentRarities = config.New("configs", "enchantment_rarities.json", defaultEnchantmentRarities)

// EnchantmentRarityName returns the enchantment rarity name
func EnchantmentRarityName(key string) string {
	data := enchantmentRarities.GetObject(key)
	return data.GetString("name")
}

// EnchantmentRarityCost returns the enchantment rarity cost
func EnchantmentRarityCost(key string) int {
	data := enchantmentRarities.GetObject(key)
	return data.GetInt("cost")
}

// EnchantmentRarityWeight returns the enchantment rarity weight
func EnchantmentRarityWeight(key string) int {
	data := enchantmentRarities.GetObject(key)
	return data.GetInt("weight")
}
