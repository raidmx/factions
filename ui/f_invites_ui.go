package ui

import (
	"fmt"

	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/player/form"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions"
	"github.com/stcraft/factions/memory"
)

type FInvitesUI struct{}

var invitesCache = map[string]map[string]string{}
var factionCache = map[string]string{}

// NewFInvitesUI
func NewFInvitesUI(player *player.Player, invites map[string]string) form.Form {
	data := config.GetUI("f_invites_menu")

	f := form.NewMenu(FInvitesUI{}, data["title"])
	f = f.WithBody(data["desc"])

	buttons := []form.Button{}

	for fac := range invites {
		buttons = append(buttons, form.NewButton(fmt.Sprintf(data["invite_title"].(string), fac), ""))
	}

	invitesCache[player.XUID()] = invites

	return f.WithButtons(
		buttons...,
	)
}

// Submit ...
func (f FInvitesUI) Submit(submitter form.Submitter, pressed form.Button) {
	p := submitter.(*player.Player)

	if pressed.Text == "" {
		delete(invitesCache, p.XUID())
		return
	}

	p.SendForm(newInvites(p, pressed.Text))
}

type Invites struct{}

// newInvites ...
func newInvites(player *player.Player, faction string) form.Form {
	data := config.GetUI("f_invites_menu")

	f := form.NewMenu(Invites{}, data["title"])
	invites := invitesCache[player.XUID()]

	f = f.WithBody(fmt.Sprintf(data["invite_desc"].(string), faction, factions.InviteExpiry(player.XUID(), faction), invites[faction]))

	factionCache[player.XUID()] = faction

	return f.WithButtons(
		form.NewButton("Accept Invite", ""), form.NewButton("§4Reject Invite", ""),
	)
}

// Submit ...
func (f Invites) Submit(submitter form.Submitter, pressed form.Button) {
	p := submitter.(*player.Player)
	fPlayer := memory.FPlayer(p)

	switch pressed.Text {
	case "Accept Invite":
		if fPlayer.Faction != nil {
			p.Message(config.Message("must_leave_faction", fPlayer.Faction.Name))
			return
		}

		// join the faction
		faction := memory.Faction(factionCache[p.XUID()])

		faction.Broadcast(config.Message("broadcast_member_joined_faction", p.Name()))
		p.Message(config.Message("joined_faction", factionCache[p.XUID()]))

		faction.AddMember(p)
		fPlayer.JoinFaction(faction)

		// delete the invite
		factions.DeleteInvite(p.XUID(), faction.Name)
		return
	case "§4Reject Invite":
		fac := factionCache[p.XUID()]
		p.Message(config.Message("rejected_invite", fac))

		// delete the invite
		factions.DeleteInvite(p.XUID(), fac)

		// broadcast
		faction := memory.Faction(fac)
		faction.BroadcastManagers(config.Message("broadcast_player_rejected_invite", p.Name()))

		delete(factionCache, fac)
		return
	}

	p.SendForm(NewFInvitesUI(p, invitesCache[p.XUID()]))
}
