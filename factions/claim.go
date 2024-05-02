package factions

import (
	"github.com/stcraft/dragonfly/server/world"
)

type Claim struct {
	Position *world.ChunkPos
	Owner    string
	Created  int64
}
