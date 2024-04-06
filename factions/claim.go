package factions

import (
	"github.com/STCraft/dragonfly/server/world"
)

type Claim struct {
	Position *world.ChunkPos
	Owner    string
	Created  int64
}
