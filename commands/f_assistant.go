package commands

import (
	"strings"

	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/loader/dragonfly"
)

type FAssistantCmd struct {
	Assistant cmd.SubCommand `cmd:"assistant"`

	Member string `cmd:"member"`
}

func (c FAssistantCmd) Run(src cmd.Source, o *cmd.Output) {
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
	newRank := factions.Assistant

	// check if has permission
	if !config.RankHasPermission(rank, "assistant") {
		mustBeRank := config.RankWithNativePermission("assistant")
		p.Message(config.Message("must_be_" + mustBeRank))

		return
	}

	// check target
	user := dragonfly.UserFromName(c.Member)

	if user == nil {
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
