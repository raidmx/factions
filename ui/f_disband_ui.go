package ui

import (
	"fmt"

	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/player/form"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions"
	"github.com/stcraft/factions/factions/chat"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/loader/dragonfly"
)

type FDisbandUI struct {
	Input form.Input
}

// NewDisbandUI ...
func NewFDisbandUI(faction string) form.Form {
	data := config.GetUI("f_disband_ui")
	input := data["input"].(map[string]any)

	return form.New(
		FDisbandUI{
			Input: form.Input{
				Text:        fmt.Sprintf(input["text"].(string), faction),
				Default:     input["default"].(string),
				Placeholder: input["placeholder"].(string),
			},
		},
		data["title"].(string),
	)
}

// Submit ...
func (f FDisbandUI) Submit(submitter form.Submitter) {
	p := submitter.(*player.Player)
	name := f.Input.Value()

	fPlayer := memory.FPlayer(p)
	if fPlayer.Faction.Name != name {
		p.Message(config.Message("faction_name_does_not_match"))
		return
	}

	p.Message(config.Message("faction_disbanded", name))
	dragonfly.Server.Broadcast(config.Message("broadcast_faction_disbanded", p.Name(), name))

	for _, m := range fPlayer.Faction.Members {
		player, ok := dragonfly.Server.PlayerByXUID(m.Xuid)

		if !ok {
			continue
		}

		if m.Rank != factions.Leader {
			player.Message(config.Message("broadcast_faction_members_disbanded"))
		}

		fPlayer := memory.FPlayer(player)
		fPlayer.LeaveFaction()

		if fPlayer.Channel.ChannelType() != chat.Global {
			fPlayer.SetChannel(chat.GlobalChannel{})
		}
	}

	memory.DeleteFaction(name)
}
