package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions"
	"github.com/stcraft/factions/ui"
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
