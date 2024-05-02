package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/factions/ui"
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
