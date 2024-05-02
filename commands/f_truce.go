package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/loader/dragonfly"
)

type FTruceCmd struct {
	Truce cmd.SubCommand `cmd:"truce"`

	Target string `cmd:"target"`
}

func (c FTruceCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)
	fMember := fPlayer.GetFMember()

	if fPlayer.Faction == nil {
		p.Message(config.Message("must_be_in_a_faction"))
		return
	}

	faction := fPlayer.Faction

	if !config.RankHasPermission(config.RankID(fMember.Rank), "truce") {
		mustBeRank := config.RankWithNativePermission("truce")
		p.Message(config.Message("must_be_" + mustBeRank))
		return
	}

	// check if target provided is player
	tPlayer, isPlayer := dragonfly.Server.PlayerByName(c.Target)
	var targetFaction *factions.Faction

	if isPlayer {
		if tPlayer.XUID() == p.XUID() {
			p.Message(config.Message("command_usage_on_self"))
			return
		}

		targetFPlayer := memory.FPlayer(tPlayer)

		if targetFPlayer.Faction == nil {
			p.Message(config.Message("player_not_in_any_faction", tPlayer.Name()))
			return
		}

		targetFaction = targetFPlayer.Faction
	} else {
		targetFaction = memory.Faction(c.Target)
	}

	if targetFaction == nil {
		p.Message(config.Message("invalid_faction_or_player", c.Target))
		return
	}

	if targetFaction.Name == faction.Name {
		p.Message(config.Message("command_usage_on_own_faction"))
		return
	}

	if targetFaction.Truce(faction) {
		p.Message(config.Message("already_truce", targetFaction.Name))
		return
	}

	if faction.MarkedTruce(targetFaction) {
		p.Message(config.Message("already_marked_truce", targetFaction.Name))
		return
	}

	faction.MarkTruce(targetFaction)

	if faction.Truce(targetFaction) {
		faction.Broadcast(config.Message("truce_established", targetFaction.Name))
		targetFaction.Broadcast(config.Message("truce_established", faction.Name))
		return
	}

	faction.Broadcast(config.Message("broadcast_marked_truce", targetFaction.Name))
	targetFaction.Broadcast(config.Message("broadcast_marked_truce_target", faction.Name, faction.Name))
}
