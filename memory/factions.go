package memory

import (
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/factions/config"
	"github.com/stcraft/factions/factions"
	"github.com/stcraft/factions/factions/warp"
	"github.com/stcraft/factions/postgres"
	"github.com/stcraft/loader/dragonfly"
)

var Factions = map[string]*factions.Faction{}

// LoadFaction loads a Faction data from the database into the memory
func LoadFaction(name string) {
	Factions[name] = postgres.Faction(name)
}

// NewFaction creates and stores the new Faction Data into the memory
func NewFaction(faction string, Creator *player.Player) {
	leader := &factions.FMember{
		Name: Creator.Name(),
		Xuid: Creator.XUID(),
		Rank: factions.Leader,
	}

	fac := &factions.Faction{
		Name:        faction,
		Description: config.GetFactionConfig[string]("default_description"),
		Allies:      []string{},
		Truces:      []string{},
		Enemies:     []string{},
		Leader:      leader,
		Members:     []*factions.FMember{leader},
		Warps:       map[string]*warp.Warp{},
		Storage:     factions.NewStorage(),
	}

	FPlayer(Creator).Faction = fac
	Factions[faction] = fac
	postgres.SaveFaction(fac)
}

// FactionExists returns whether a Faction with a particular name exists in the memory
func FactionExists(faction string) bool {
	return Factions[faction] != nil
}

// Faction gets a faction data from the memory if loaded otherwise loads it
func Faction(faction string) *factions.Faction {
	data, loaded := Factions[faction]

	if !loaded {
		LoadFaction(faction)
		return Factions[faction]
	}

	return data
}

// DeleteFaction deletes the faction data both from the memory & the database
func DeleteFaction(faction string) {
	for _, m := range Faction(faction).Members {
		p, ok := dragonfly.Server.PlayerByXUID(m.Xuid)

		if !ok {
			return
		}

		FPlayer(p).Faction = nil
	}

	delete(Factions, faction)
	postgres.DeleteFaction(faction)
}

// SaveFaction saves the Faction data from the memory into the database
func SaveFaction(faction *factions.Faction) {
	postgres.UpdateFaction(faction)
}

// SaveAllFactions saves all the Faction data from the memory into the database
func SaveAllFactions() {
	for _, f := range Factions {
		SaveFaction(f)
	}
}
