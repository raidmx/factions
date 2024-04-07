package commands

import (
	"time"

	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/factions/home"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FSetHomeCmd struct {
	SetHome cmd.SubCommand `cmd:"sethome"`
}

func (FSetHomeCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	if fPlayer.Faction == nil {
		p.Message(config.Message("must_be_in_a_faction"))
		return
	}

	fMember := fPlayer.GetFMember()
	faction := fPlayer.Faction

	if !config.RankHasPermission(config.RankID(fMember.Rank), "sethome") {
		mustBeRank := config.RankWithNativePermission("sethome")
		p.Message(config.Message("must_be_" + mustBeRank))
		return
	}

	// Set the current location as the faction home
	owner := memory.ChunkOwner(memory.ChunkPos(p.Position(), p.World()))

	// prevent setting home in chunk claimed by other faction
	if owner != nil && owner.Name != faction.Name {
		p.Message(config.Message("cannot_set_home", owner.Name))
		return
	}

	faction.Home = home.New(p.Position(), p.Rotation(), time.Now().Unix(), p.Name())
	p.Message(config.Message("faction_home_set"))
}
