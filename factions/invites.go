package factions

import (
	"time"

	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/Factions/config"
)

// Invite sets the invite key for the player's XUID for the expiry duration
func Invite(xuid string, faction string, invitedBy string) {
	dur := int64(config.GetFactionConfig[float64]("invitation_expiry"))
	expiry := time.Now().Unix() + dur

	dragonfly.DBExec(`INSERT INTO "FINVITES" ("XUID", "FACTION", "INVITED_BY", "EXPIRY") VALUES ($1, $2, $3, $4)`, xuid, faction, invitedBy, expiry)
}

// AllInvites returns a map of all invites of the specified player with the following data:
// Key: Faction name
// Value: Player from the faction that invited
func AllInvites(xuid string) map[string]string {
	invites := map[string]string{}
	rows := dragonfly.DBQuery(`SELECT "FACTION", "INVITED_BY" FROM "FINVITES" WHERE "XUID" = $1`, xuid)
	defer rows.Close()

	for rows.Next() {
		var faction, invitedBy string
		rows.Scan(&faction, &invitedBy)

		invites[faction] = invitedBy
	}

	return invites
}

// InviteExpiry returns the expiry of a faction invite
func InviteExpiry(xuid string, faction string) time.Duration {
	rows := dragonfly.DBQuery(`SELECT "EXPIRY" FROM "FINVITES" WHERE "XUID" = $1 AND "FACTION" = $2`, xuid, faction)
	defer rows.Close()

	if !rows.Next() {
		return time.Second * 0
	}

	var expiry int64
	rows.Scan(&expiry)

	return time.Duration(expiry)
}

// CheckInvite returns whether the faction has already invited this player
func CheckInvite(xuid string, faction string) bool {
	invites := AllInvites(xuid)

	for f := range invites {
		if f == faction {
			return true
		}
	}

	return false
}

// DeleteInvite deletes the faction invite
func DeleteInvite(xuid string, faction string) {
	dragonfly.DBExec(`DELETE FROM "FINVITES" WHERE "XUID" = $1 AND "FACTION" = $2`, xuid, faction)
}
