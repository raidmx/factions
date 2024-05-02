package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/factions/ui"
)

type FCreateCmd struct {
	Create cmd.SubCommand `cmd:"create"`
}

func (FCreateCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	if fPlayer.Faction != nil {
		faction := fPlayer.Faction

		if faction.Leader.Xuid == p.XUID() {
			p.Message(config.Message("must_disband_faction", faction.Name))
			return
		}

		p.Message(config.Message("must_leave_faction", faction.Name))
		return
	}

	p.SendForm(ui.NewFCreateUI())
}
