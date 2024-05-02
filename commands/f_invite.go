package commands

import (
	"fmt"

	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions"
	"github.com/stcraft/factions/memory"
	"github.com/stcraft/factions/ui"
)

type FInviteCmd struct {
	Invite cmd.SubCommand `cmd:"invite"`

	Target cmd.Optional[[]cmd.Target] `cmd:"player"`
}

func (c FInviteCmd) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)

	if !ok {
		o.Print(config.Message("command_usage_by_console"))
		return
	}

	fPlayer := memory.FPlayer(p)

	if fPlayer.Faction == nil {
		p.Message(config.Message("must_be_in_a_faction"))
		return
	} else if fPlayer.GetFMember().Rank < factions.Manager {
		p.Message(config.Message("must_be_manager"))
		return
	}

	faction := fPlayer.Faction

	maxMembers := int(config.GetFactionConfig[float64]("total_faction_members"))
	if maxMembers == len(faction.Members) {
		p.Message(config.Message("max_members", fmt.Sprint(maxMembers)))
		return
	}

	targets, _ := c.Target.Load()

	if len(targets) == 0 {
		p.SendForm(ui.NewFInviteUI())
		return
	}

	if len(targets) > 1 {
		p.Message(config.Message("more_than_one_target"))
		return
	}

	target := targets[0].(*player.Player)

	if target == p {
		p.Message(config.Message("command_usage_on_self"))
		return
	}

	// check if target is in same faction
	if faction.IsMember(target) {
		p.Message(config.Message("already_a_member", target.Name()))
		return
	}

	targetFPlayer := memory.FPlayer(target)
	targetFPlayer.Invite(faction, p)
}
