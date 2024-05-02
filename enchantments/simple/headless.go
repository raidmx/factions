package simple

import (
	"github.com/stcraft/dragonfly/server/item"
	"github.com/stcraft/dragonfly/server/world"
	"github.com/stcraft/factions/config"
)

type Headless struct{}

func (Headless) Name() string {
	return config.EnchantmentName("headless")
}

func (Headless) MaxLevel() int {
	return config.MaxEnchantmentLevel("headless")
}

func (Headless) Cost(level int) (int, int) {
	return config.MinEnchantmentCost("headless"), config.MaxEnchantmentCost("headless")
}

func (Headless) Rarity() item.EnchantmentRarity {
	return SimpleRarity{}
}

func (Headless) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return true
}

func (Headless) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Sword)
	return ok
}
