package commands

import (
	"fmt"

	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions/board"
	"github.com/stcraft/factions/memory"
)

type FMapCmd struct {
	Map cmd.SubCommand `cmd:"map"`

	AutoUpdate cmd.Optional[bool] `cmd:"autoupdate"`
}

func (c FMapCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	autoUpdate, ok := c.AutoUpdate.Load()
	if !ok {
		autoUpdate = false
	}

	if fPlayer.AutoUpdate != autoUpdate {
		p.Message(config.Message("faction_map_autoupdate", fmt.Sprint(autoUpdate)))
		fPlayer.AutoUpdate = autoUpdate
		return
	}

	// send the board
	p.Message(board.FactionMap(fPlayer))
}
