package commands

import (
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FDescCmd struct {
	Desc cmd.SubCommand `cmd:"desc"`

	Description string `cmd:"description"`
}

func (c FDescCmd) Run(src cmd.Source, o *cmd.Output) {
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

	faction := fPlayer.Faction
	fMember := fPlayer.GetFMember()
	rank := config.RankID(fMember.Rank)

	// check if has permission
	if !config.RankHasPermission(rank, "desc") {
		mustBeRank := config.RankWithNativePermission("desc")
		p.Message(config.Message("must_be_" + mustBeRank))

		return
	}

	// change the description
	faction.Description = c.Description
	p.Message(config.Message("description_changed", faction.Description))
}
