package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/memory"
)

type FHomeCmd struct {
	Home cmd.SubCommand `cmd:"home"`
}

func (FHomeCmd) Run(src cmd.Source, o *cmd.Output) {
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

	if !config.RankHasPermission(config.RankID(fMember.Rank), "home") {
		mustBeRank := config.RankWithNativePermission("home")
		p.Message(config.Message("must_be_" + mustBeRank))
		return
	}

	if faction.Home == nil {
		p.Message(config.Message("faction_has_no_home"))
		return
	}

	// teleport to the faction home
	faction.Home.Teleport(p)
}
