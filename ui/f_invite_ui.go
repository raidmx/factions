package ui

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/player/form"
)

type FInviteUI struct {
	Input form.Input
}

// NewFInviteUI ...
func NewFInviteUI() form.Form {
	data := config.GetUI("f_invite_ui")
	input := data["input"].(map[string]any)

	return form.New(
		FInviteUI{
			Input: form.Input{
				Text:        input["text"].(string),
				Default:     input["default"].(string),
				Placeholder: input["placeholder"].(string),
			},
		},
		data["title"].(string),
	)
}

func (f FInviteUI) Submit(submitter form.Submitter) {
	p := submitter.(*player.Player)
	fPlayer := memory.FPlayer(p)

	name := f.Input.Value()

	if len(name) < 3 || len(name) > 12 {
		p.Message(config.Message("player_name_invalid"))
		return
	}

	target, ok := dragonfly.Server.PlayerByName(name)

	if !ok {
		p.Message(config.Message("player_not_found", name))
		return
	}

	if target == p {
		p.Message(config.Message("command_usage_on_self"))
		return
	}

	targetFPlayer := memory.FPlayer(target)
	targetFPlayer.Invite(fPlayer.Faction, p)
}
