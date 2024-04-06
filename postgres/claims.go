package postgres

import (
	"encoding/json"

	"github.com/STCraft/DFLoader/db"
	"github.com/STCraft/dragonfly/server/world"
	"github.com/inceptionmc/factions/factions"
)

// GetAllClaims returns all the claims from the database
func GetAllClaims() map[world.ChunkPos]*factions.Claim {
	claims := map[world.ChunkPos]*factions.Claim{}

	rows := db.DB.Query(`SELECT * FROM CLAIMS`)

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
	db.DB.Exec(`INSERT INTO CLAIMS(POSITION, OWNER, CREATED) VALUES($1, $2, $3)`, position, owner, created)
}

// DeleteClaim deletes the claim at a location
func DeleteClaim(pos string) {
	db.DB.Exec(`DELETE FROM CLAIMS WHERE POSITION = $1`, pos)
}
