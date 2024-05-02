package ui

import (
	"unicode"

	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/player/form"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/factions/postgres"
	"github.com/stcraft/loader/dragonfly"
)

type FCreateUI struct {
	Input form.Input
}

func NewFCreateUI() form.Form {
	data := config.GetUI("f_create_ui")
	input := data["input"].(map[string]any)

	return form.New(
		FCreateUI{
			Input: form.Input{
				Text:        input["text"].(string),
				Default:     input["default"].(string),
				Placeholder: input["placeholder"].(string),
			},
		},
		data["title"].(string),
	)
}

// Submit ...
func (f FCreateUI) Submit(submitter form.Submitter) {
	p := submitter.(*player.Player)
	name := f.Input.Value()

	// check if faction with the name exists already
	if memory.FactionExists(name) || postgres.FactionExists(name) {
		p.Message(config.Message("faction_name_exists", name))
		return
	}

	if len(name) < 3 {
		p.Message(config.Message("faction_name_too_small"))
		return
	}

	if len(name) > 15 {
		p.Message(config.Message("faction_name_too_long"))
		return
	}

	if !ValidFactionName(name) {
		p.Message(config.Message("faction_name_invalid"))
		return
	}

	memory.NewFaction(name, p)

	p.Message(config.Message("faction_created", name))
	dragonfly.Server.Broadcast(config.Message("broadcast_faction_created", p.Name(), name))
}

// ValidFactionName returns whether the faction name is valid
func ValidFactionName(name string) bool {
	for _, c := range name {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			valid := false

			for _, a := range []rune{'#', '@', '$', '_'} {
				if c == a {
					valid = true
					break
				}
			}

			if !valid {
				return valid
			}
		}
	}

	return true
}
