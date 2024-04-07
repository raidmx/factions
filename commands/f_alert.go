package commands

import (
	"time"

	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

type FAlertCmd struct {
	Alert cmd.SubCommand `cmd:"alert"`

	Type AlertType `cmd:"type"`

	Message string `cmd:"message"`
}

func (c FAlertCmd) Run(src cmd.Source, o *cmd.Output) {
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
	if !config.RankHasPermission(rank, "alert") {
		mustBeRank := config.RankWithNativePermission("alert")
		p.Message(config.Message("must_be_" + mustBeRank))

		return
	}

	switch c.Type {
	case "title":
		titleData := config.TitleData("faction_alert")
		fadeIn := time.Duration(titleData["fadeIn"].(float64)) * time.Second
		fadeOut := time.Duration(titleData["fadeOut"].(float64)) * time.Second
		stay := time.Duration(titleData["stay"].(float64)) * time.Second

		faction.BroadcastTitle(titleData["title"].(string), c.Message, fadeIn, fadeOut, stay)
	case "popup":
		faction.BroadcastPopup("§cFaction Alert: §f" + c.Message)
	case "notification":
		faction.BroadcastToast("§cFaction Alert §8[§7"+p.Name()+"§8]", c.Message)
	}
}

type AlertType string

func (AlertType) Type() string {
	return "type"
}

func (AlertType) Options(_ cmd.Source) []string {
	return []string{"title", "popup", "notification"}
}
