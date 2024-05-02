package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/loader/dragonfly"
)

type FEnemyCmd struct {
	Enemy cmd.SubCommand `cmd:"enemy"`

	Target string `cmd:"target"`
}

func (c FEnemyCmd) Run(src cmd.Source, o *cmd.Output) {
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

	if !config.RankHasPermission(config.RankID(fMember.Rank), "enemy") {
		mustBeRank := config.RankWithNativePermission("enemy")
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

	if targetFaction.Enemy(faction) {
		p.Message(config.Message("already_enemy", targetFaction.Name))
		return
	}

	faction.MarkEnemy(targetFaction)
	targetFaction.MarkEnemy(faction)

	faction.Broadcast(config.Message("broadcast_marked_enemy", targetFaction.Name))
	targetFaction.Broadcast(config.Message("broadcast_marked_enemy_target", faction.Name))
}
