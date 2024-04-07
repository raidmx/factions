package postgres

import (
	"encoding/json"

	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/Factions/factions"
	"github.com/STCraft/Factions/utils"
	"github.com/STCraft/dragonfly/server/world"
)

// GetAllClaims returns all the claims from the database
func GetAllClaims() map[world.ChunkPos]*factions.Claim {
	claims := map[world.ChunkPos]*factions.Claim{}

	rows := dragonfly.DBQuery(`SELECT * FROM "CLAIMS"`)

	for rows.Next() {
		var p, owner string
		var created int64
		rows.Scan(&p, &owner, &created)

		var position *world.ChunkPos

		json.Unmarshal([]byte(p), &position)

		claims[*position] = &factions.Claim{
			Position: position,
			Owner:    owner,
			Created:  created,
		}
	}

	return claims
}

// RegisterClaim registers a new claim
func RegisterClaim(position *world.ChunkPos, owner string, created int64) {
	dragonfly.DBExec(`INSERT INTO "CLAIMS" ("POSITION", "OWNER", "CREATED") VALUES($1, $2, $3)`, utils.Encode(position), owner, created)
}

// DeleteClaim deletes the claim at a location
func DeleteClaim(pos string) {
	dragonfly.DBExec(`DELETE FROM "CLAIMS" WHERE "POSITION" = $1`, pos)
}
