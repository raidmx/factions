package commands

import (
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/factions"
	"github.com/STCraft/Factions/ui"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FInvitesCmd struct {
	Invite cmd.SubCommand `cmd:"invites"`
}

func (FInvitesCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	invites := factions.AllInvites(p.XUID())

	if len(invites) == 0 {
		p.Message(config.Message("zero_invites"))
		return
	}

	p.SendForm(ui.NewFInvitesUI(p, invites))
}
