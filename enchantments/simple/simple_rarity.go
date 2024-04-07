package simple

import "github.com/STCraft/Factions/config"

type SimpleRarity struct{}

func (SimpleRarity) Name() string {
	return config.EnchantmentRarityName("simple")
}

func (SimpleRarity) Cost() int {
	return config.EnchantmentRarityCost("simple")
}

func (SimpleRarity) Weight() int {
	return config.EnchantmentRarityWeight("simple")
}
