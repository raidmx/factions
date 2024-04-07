package commands

import (
	"fmt"
	"strings"

	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/factions"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FManagerCmd struct {
	Manager cmd.SubCommand `cmd:"manager"`

	Member string `cmd:"member"`
}

func (c FManagerCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	// check if console
	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	// check if faction exists
	if fPlayer.Faction == nil {
		p.Message(config.Message("must_be_in_a_faction"))
		return
	}

	faction := fPlayer.Faction
	fMember := fPlayer.GetFMember()
	rank := config.RankID(fMember.Rank)
	newRank := factions.Manager

	// check if has permission
	if !config.RankHasPermission(rank, "manager") {
		mustBeRank := config.RankWithNativePermission("manager")
		p.Message(config.Message("must_be_" + mustBeRank))

		return
	}

	// check limits
	coMax := int(config.GetFactionConfig[float64]("number_of_managers"))
	if len(faction.MembersWithRank(newRank))+1 == coMax {
		p.Message(config.Message("max_managers", fmt.Sprint(coMax)))
		return
	}

	// check target
	user := dragonfly.UserFromName(c.Member)

	if !ok {
		p.Message(config.Message("invalid_player"))
		return
	}

	if strings.EqualFold(user.Name, p.Name()) {
		p.Message(config.Message("command_usage_on_self"))
		return
	}

	targetFMember := faction.TryGetMember(c.Member)

	if targetFMember == nil {
		p.Message(config.Message("player_not_in_faction", user.Name))
		return
	}

	// set the rank
	targetFMember.Rank = newRank
	rank = config.RankName(newRank)

	p.Message(config.Message("faction_rank_changed", user.Name, rank))
	faction.Broadcast(config.Message("broadcast_faction_rank_changed", p.Name(), user.Name, rank))
}
