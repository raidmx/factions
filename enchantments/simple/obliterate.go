package simple

import (
	"github.com/stcraft/dragonfly/server/item"
	"github.com/stcraft/dragonfly/server/world"
	"github.com/stcraft/factions/config"
)

type Obliterate struct{}

func (Obliterate) Name() string {
	return config.EnchantmentName("obliterate")
}

func (Obliterate) MaxLevel() int {
	return config.MaxEnchantmentLevel("obliterate")
}

func (Obliterate) Cost(level int) (int, int) {
	return config.MinEnchantmentCost("obliterate"), config.MaxEnchantmentCost("obliterate")
}

func (Obliterate) Rarity() item.EnchantmentRarity {
	return SimpleRarity{}
}

func (Obliterate) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	return true
}

func (Obliterate) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Sword)
	return ok
}
