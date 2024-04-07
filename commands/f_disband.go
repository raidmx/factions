package commands

import (
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/Factions/ui"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FDisbandCmd struct {
	Disband cmd.SubCommand `cmd:"disband"`
}

func (FDisbandCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	if !fPlayer.IsInAnyFaction() {
		p.Message(config.Message("must_be_in_a_faction"))
		return
	} else if fPlayer.Faction.Leader.Xuid != p.XUID() {
		p.Message(config.Message("must_be_leader"))
		return
	}

	p.SendForm(ui.NewFDisbandUI(fPlayer.Faction.Name))
}
