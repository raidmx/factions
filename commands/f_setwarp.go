package commands

import (
	"time"

	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/factions/warp"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FSetWarpCmd struct {
	SetWarp cmd.SubCommand `cmd:"setwarp"`

	Name string `cmd:"name"`
}

func (c FSetWarpCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	if fPlayer.Faction == nil {
		p.Message(config.Message("must_be_in_a_faction"))
		return
	}

	fMember := fPlayer.GetFMember()
	faction := fPlayer.Faction

	if !config.RankHasPermission(config.RankID(fMember.Rank), "setwarp") {
		mustBeRank := config.RankWithNativePermission("setwarp")
		p.Message(config.Message("must_be_" + mustBeRank))
		return
	}

	// Check if warp exists
	if faction.WarpExists(c.Name) {
		p.Message(config.Message("faction_warp_exists", c.Name))
		return
	}

	owner := memory.ChunkOwner(memory.ChunkPos(p.Position(), p.World()))

	// prevent setting home in chunk claimed by other faction
	if owner != nil && owner.Name != faction.Name {
		p.Message(config.Message("cannot_set_warp", owner.Name))
		return
	}

	// Set the current location as the faction warp
	faction.Warps[c.Name] = warp.New(c.Name, p.Position(), p.Rotation(), time.Now().Unix(), p.Name())
	p.Message(config.Message("faction_warp_set", c.Name))
}
