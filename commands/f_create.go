package commands

import (
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/Factions/ui"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
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
