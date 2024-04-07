package commands

import (
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FUnclaimAllCmd struct {
	UnclaimAll cmd.SubCommand `cmd:"unclaimall"`

	Confirmation cmd.Optional[string] `cmd:"confirmation"`
}

func (c FUnclaimAllCmd) Run(src cmd.Source, o *cmd.Output) {
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
	if !config.RankHasPermission(rank, "unclaimall") {
		mustBeRank := config.RankWithNativePermission("unclaimall")
		p.Message(config.Message("must_be_" + mustBeRank))

		return
	}

	// confirmation
	conf, ok := c.Confirmation.Load()

	if !ok {
		p.Message(config.Message("confirmation_error", "f unclaimall", faction.Name))
		return
	}

	if conf != faction.Name {
		p.Message(config.Message("confirmation_does_not_match"))
		return
	}

	// unclaim all the claims
	claims := memory.GetClaims(faction.Name)
	count := len(claims)

	for _, c := range claims {
		memory.DeleteClaim(c.Position)
	}

	p.Message(config.Message("all_chunks_unclaimed", count))
}
