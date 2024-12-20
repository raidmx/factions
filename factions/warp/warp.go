package warp

import (
	"github.com/stcraft/dragonfly/server/block/cube"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/factions/teleport"
)

type Warp struct {
	Name string

	Location [3]float64

	Yaw float64

	Pitch float64

	Created int64

	CreatedBy string
}

func New(name string, loc [3]float64, rotation cube.Rotation, created int64, createdBy string) *Warp {
	return &Warp{
		Name:      name,
		Location:  loc,
		Yaw:       rotation.Yaw(),
		Pitch:     rotation.Pitch(),
		Created:   created,
		CreatedBy: createdBy,
	}
}

// Teleport teleports the player to the warp
func (w *Warp) Teleport(p *player.Player) {
	teleport.Teleport(p, w.Location, w.Yaw, w.Pitch)
}
