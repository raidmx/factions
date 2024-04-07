package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/factions"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FAllyCmd struct {
	Ally cmd.SubCommand `cmd:"ally"`

	Target string `cmd:"target"`
}

func (c FAllyCmd) Run(src cmd.Source, o *cmd.Output) {
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

	if !config.RankHasPermission(config.RankID(fMember.Rank), "ally") {
		mustBeRank := config.RankWithNativePermission("ally")
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

	if targetFaction.Alliance(faction) {
		p.Message(config.Message("already_ally", targetFaction.Name))
		return
	}

	if faction.MarkedAlly(targetFaction) {
		p.Message(config.Message("already_marked_ally", targetFaction.Name))
		return
	}

	faction.MarkAlly(targetFaction)

	if faction.Alliance(targetFaction) {
		faction.Broadcast(config.Message("ally_established", targetFaction.Name))
		targetFaction.Broadcast(config.Message("ally_established", faction.Name))
		return
	}

	faction.Broadcast(config.Message("broadcast_marked_ally", targetFaction.Name))
	targetFaction.Broadcast(config.Message("broadcast_marked_ally_target", faction.Name, faction.Name))
}
