package commands

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/factions/ui"
)

type FInfoCmd struct {
	Info cmd.SubCommand `cmd:"info"`

	Faction cmd.Optional[string] `cmd:"faction"`
}

func (c FInfoCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fac, ok := c.Faction.Load()
	fPlayer := memory.FPlayer(p)
	faction := memory.Faction(fac)

	if !ok {
		if fPlayer.Faction == nil {
			p.Message(config.Message("faction_not_found"))
			return
		}

		faction = fPlayer.Faction
	}

	if faction == nil {
		p.Message(config.Message("faction_not_found"))
		return
	}

	p.SendForm(ui.NewFInfoUI(p, faction))
}
