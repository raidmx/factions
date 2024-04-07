package postgres

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/Factions/config"
	"github.com/STCraft/Factions/factions"
	"github.com/STCraft/Factions/factions/chat"
	"github.com/STCraft/dragonfly/server/player"
)

// FPlayer returns the data from the database
func FPlayer(p *player.Player) (fPlayer *factions.FPlayer, faction string) {
	rows := dragonfly.DBQuery(`SELECT "FACTION", "CHANNEL" FROM "FPLAYERS" WHERE "XUID" = $1`, p.XUID())
	defer rows.Close()

	if rows.Next() {
		var fac, channel string
		rows.Scan(&fac, &channel)

		fPlayer := &factions.FPlayer{
			Player:  p,
			Faction: nil,
			Channel: config.ChannelFromID(channel),
		}
		return fPlayer, fac
	}

	return nil, ""
}

// SaveFPlayer saves the player data into the Database
func SaveFPlayer(fPlayer *factions.FPlayer) {
	faction := fPlayer.Faction
	var fName string

	if faction != nil {
		fName = faction.Name
	}

	dragonfly.DBExec(`INSERT INTO "FPLAYERS" ("XUID", "FACTION", "CHANNEL") VALUES($1, $2, $3)`, fPlayer.Player.XUID(), fName, chat.ChannelID(fPlayer.Channel))
}

// UpdateFPlayer updates the player data into the Database
func UpdateFPlayer(fPlayer *factions.FPlayer) {
	faction := fPlayer.Faction

	fName := ""
	if faction != nil {
		fName = faction.Name
	}

	dragonfly.DBExec(`UPDATE "FPLAYERS" SET "FACTION" = $1, "CHANNEL" = $2 where "XUID" = $3`, fName, chat.ChannelID(fPlayer.Channel), fPlayer.Player.XUID())
}

// FPlayerExists returns whether the player data exists for a player
func FPlayerExists(xuid string) bool {
	rows := dragonfly.DBQuery(`SELECT * FROM "FPLAYERS" WHERE "XUID" = $1`, xuid)
	defer rows.Close()
	return rows.Next()
}
