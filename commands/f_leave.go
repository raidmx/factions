package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions/chat"
	"github.com/stcraft/factions/memory"
)

type FLeaveCmd struct {
	Leave cmd.SubCommand `cmd:"leave"`
}

func (FLeaveCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	if !fPlayer.IsInAnyFaction() {
		p.Message(config.Message("must_be_in_a_faction"))
		return
	} else if fPlayer.Faction.Leader.Xuid == p.XUID() {
		p.Message(config.Message("leader_cannot_leave"))
		return
	}

	// leave the faction
	faction := fPlayer.Faction

	p.Message(config.Message("left_faction", faction.Name))
	faction.Broadcast(config.Message("broadcast_member_left_faction", p.Name(), config.RankName(fPlayer.GetFMember().Rank)))

	fPlayer.LeaveFaction()

	if fPlayer.Channel.ChannelType() != chat.Global {
		fPlayer.SetChannel(chat.GlobalChannel{})
	}
}
