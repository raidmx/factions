package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/factions"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/player/title"
)

type FDemoteCmd struct {
	Demote cmd.SubCommand `cmd:"demote"`

	Member string `cmd:"member"`
}

func (c FDemoteCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	// check if console
	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	// check if faction exists
	if fPlayer.Faction == nil {
		p.Message(config.Message("must_be_in_a_faction"))
		return
	}

	faction := fPlayer.Faction
	fMember := fPlayer.GetFMember()
	rank := config.RankID(fMember.Rank)

	// check if has permission
	if !config.RankHasPermission(rank, "demote") {
		mustBeRank := config.RankWithNativePermission("demote")
		p.Message(config.Message("must_be_" + mustBeRank))

		return
	}

	// check target
	user := dragonfly.UserFromName(c.Member)

	if user == nil {
		p.Message(config.Message("invalid_player"))
		return
	}

	if strings.EqualFold(user.Name, p.Name()) {
		p.Message(config.Message("command_usage_on_self"))
		return
	}

	targetFMember := faction.TryGetMember(c.Member)

	if targetFMember == nil {
		p.Message(config.Message("player_not_in_faction", user.Name))
		return
	}

	// check if higher in hierarchy
	if !fMember.Compare(targetFMember) {
		p.Message(config.Message("must_be_higher_in_hierarchy", user.Name))
		return
	}

	// check if target is already recruit
	if targetFMember.Rank == factions.Recruit {
		p.Message(config.Message("cannot_demote_recruit"))
		return
	}

	// demote
	newRank := targetFMember.Predecessor()

	officersMax := int(config.GetFactionConfig[float64]("number_of_officers"))
	managersMax := int(config.GetFactionConfig[float64]("number_of_managers"))

	switch newRank {
	case factions.Officer:
		if len(faction.MembersWithRank(newRank))+1 == officersMax {
			p.Message(config.Message("max_officers", fmt.Sprint(officersMax)))
			return
		}
	case factions.Manager:
		if len(faction.MembersWithRank(newRank))+1 == managersMax {
			p.Message(config.Message("max_managers", fmt.Sprint(managersMax)))
			return
		}
	}

	// demote
	fMember.Rank = newRank
	faction.Broadcast(config.Message("faction_member_demoted", targetFMember.Name, config.RankName(newRank)))

	// send title
	targetPlayer, ok := dragonfly.Server.PlayerByName(user.Name)

	if !ok {
		return
	}

	titleData := config.TitleData("demoted")
	fadeIn := time.Duration(titleData["fadeIn"].(float64)) * time.Second
	fadeOut := time.Duration(titleData["fadeOut"].(float64)) * time.Second
	stay := time.Duration(titleData["stay"].(float64)) * time.Second

	title := title.New(titleData["title"]).WithSubtitle(fmt.Sprintf(titleData["subtitle"].(string), config.RankName(newRank))).WithFadeInDuration(fadeIn).WithFadeOutDuration(fadeOut).WithDuration(stay)
	targetPlayer.SendTitle(title)
}
