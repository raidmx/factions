package main

import (
	"github.com/STCraft/DFLoader/loader"
	"github.com/STCraft/dragonfly/server"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/inceptionmc/factions/commands"
	"github.com/inceptionmc/factions/memory"
	"github.com/inceptionmc/factions/postgres"
)

func main() {
	registerCommands()
	loader.Init()
	postgres.Init()

	c := memory.LoadClaims()
	server.Console.Printf("[Factions] Loaded %d claims!", c)

	loader.Start()
}

func onPlayerJoin(p *player.Player) {
	srv.PlayerJoined(p)

	// Register Handlers
	p.Handle(Handler(p))

	// Load Player Data
	memory.LoadFPlayer(p)
}

// registerCommands registers all the commands
func registerCommands() {
	cmd.Register(cmd.New("faction", "Faction commands", []string{"f"}, commands.FCreateCmd{}, commands.FDescCmd{}, commands.FDisbandCmd{}, commands.FInviteCmd{}, commands.FInfoCmd{}, commands.FPromoteCmd{}, commands.FDemoteCmd{}, commands.FChatCmd{}, commands.FLeaveCmd{}, commands.FInvitesCmd{}, commands.FKickCmd{}, commands.FTruceCmd{}, commands.FAllyCmd{}, commands.FNeutralCmd{}, commands.FEnemyCmd{}, commands.FHomeCmd{}, commands.FWarpCmd{}, commands.FSetHomeCmd{}, commands.FSetWarpCmd{}, commands.FAlertCmd{}, commands.FClaimCmd{}, commands.FAutoClaimCmd{}, commands.FUnclaimCmd{}, commands.FUnclaimAllCmd{}, commands.FMapCmd{}, commands.FRecruitCmd{}, commands.FAssistantCmd{}, commands.FOfficerCmd{}, commands.FManagerCmd{}, commands.FColeaderCmd{}))
	cmd.Register(cmd.New("test", "Test command", []string{}, commands.TestCmd{}))
}
