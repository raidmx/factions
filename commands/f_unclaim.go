package commands

import (
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FUnclaimCmd struct {
	Unclaim cmd.SubCommand `cmd:"unclaim"`
}

func (FUnclaimCmd) Run(src cmd.Source, o *cmd.Output) {
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

	faction := fPlayer.Faction
	fMember := fPlayer.GetFMember()
	rank := config.RankID(fMember.Rank)

	// check if has permission
	if !config.RankHasPermission(rank, "claim") {
		mustBeRank := config.RankWithNativePermission("claim")
		p.Message(config.Message("must_be_" + mustBeRank))

		return
	}

	chunkPos := memory.ChunkPos(p.Position(), p.World())
	owner := memory.ChunkOwner(chunkPos)

	// check if land is already free
	if owner == nil {
		p.Message(config.Message("chunk_is_wilderness"))
		return
	}

	// check if land is owned by another faction
	if owner != nil && owner.Name != faction.Name {
		p.Message(config.Message("cannot_unclaim_chunk", owner.Name))
		return
	}

	// check if any warps or homes exist in this chunk
	// if they do, remove it
	if faction.Home != nil {
		chunk := memory.ChunkPos(faction.Home.Location, p.World())
		if chunk.X() == chunkPos.X() && chunk.Z() == chunkPos.Z() {
			faction.Home = nil
		}
	}

	for _, w := range faction.Warps {
		chunk := memory.ChunkPos(w.Location, p.World())
		if chunk.X() == chunkPos.X() && chunk.Z() == chunkPos.Z() {
			faction.RemoveWarp(w.Name)
		}
	}

	// unclaim
	memory.DeleteClaim(chunkPos)
	p.Message(config.Message("chunk_unclaimed", chunkPos.X(), chunkPos.Z()))
}
