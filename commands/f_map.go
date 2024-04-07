package commands

import (
	"fmt"

	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/factions/board"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
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
