package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/memory"
)

type FAutoClaimCmd struct {
	AutoClaim cmd.SubCommand `cmd:"autoclaim"`
}

func (FAutoClaimCmd) Run(src cmd.Source, o *cmd.Output) {
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
	rank := config.RankID(fMember.Rank)

	// check if has permission
	if !config.RankHasPermission(rank, "autoclaim") {
		mustBeRank := config.RankWithNativePermission("autoclaim")
		p.Message(config.Message("must_be_" + mustBeRank))

		return
	}

	fPlayer.AutoClaim = !fPlayer.AutoClaim
	p.Message(config.Message("autoclaim_changed", fPlayer.AutoClaim))
}
