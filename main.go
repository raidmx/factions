package main

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/DFLoader/loader"
	"github.com/STCraft/Factions/commands"
	"github.com/STCraft/Factions/memory"
	"github.com/STCraft/Factions/postgres"
	"github.com/STCraft/dragonfly/server"
	"github.com/STCraft/dragonfly/server/cmd"
)

func main() {
	defer func() {
		memory.SaveAllFPlayers()
		memory.SaveAllFactions()
		loader.Deinit()
	}()

	loader.Init() // Intiialise the dragonfly server

	postgres.Init()                                                 // Initialise our tables in the PostgreSQL database
	registerCommands()                                              // Register all our commands
	dragonfly.Server.RegisterHandler("factions", &FactionHandler{}) // Register Faction Handler

	c := memory.LoadClaims()
	server.Console.SendMessage("[Factions] Loaded %d claims!", c)

	loader.Start() // Start the dragonfly server
}

// registerCommands registers all the commands
func registerCommands() {
	cmd.Register(cmd.New("faction", "Faction commands", []string{"f"}, commands.FCreateCmd{}, commands.FDescCmd{}, commands.FDisbandCmd{}, commands.FInviteCmd{}, commands.FInfoCmd{}, commands.FPromoteCmd{}, commands.FDemoteCmd{}, commands.FChatCmd{}, commands.FLeaveCmd{}, commands.FInvitesCmd{}, commands.FKickCmd{}, commands.FTruceCmd{}, commands.FAllyCmd{}, commands.FNeutralCmd{}, commands.FEnemyCmd{}, commands.FHomeCmd{}, commands.FWarpCmd{}, commands.FSetHomeCmd{}, commands.FSetWarpCmd{}, commands.FAlertCmd{}, commands.FClaimCmd{}, commands.FAutoClaimCmd{}, commands.FUnclaimCmd{}, commands.FUnclaimAllCmd{}, commands.FMapCmd{}, commands.FRecruitCmd{}, commands.FAssistantCmd{}, commands.FOfficerCmd{}, commands.FManagerCmd{}, commands.FColeaderCmd{}))
}
