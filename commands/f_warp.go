package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/memory"
)

type FWarpCmd struct {
	Warp cmd.SubCommand `cmd:"warp"`

	Name string `cmd:"name"`
}

func (c FWarpCmd) Run(src cmd.Source, o *cmd.Output) {
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

	if !config.RankHasPermission(config.RankID(fMember.Rank), "warp") {
		mustBeRank := config.RankWithNativePermission("warp")
		p.Message(config.Message("must_be_" + mustBeRank))
		return
	}

	if !faction.WarpExists(c.Name) {
		p.Message(config.Message("faction_warp_does_not_exist", c.Name))
		return
	}

	// teleport to the faction warp
	faction.Warps[c.Name].Teleport(p)
}
