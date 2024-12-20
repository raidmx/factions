package config

import (
	_ "embed"

	"github.com/stcraft/factions/factions/chat"
)

// FactionConfig represents the main Faction configuration format
// which contains the major configuratives to change how the faction
// server should work.
type FactionConfig struct {
	DefaultDescription string `json:"default_description"`
	InviteExpiry       int    `json:"invite_expiry"`
	TeleportDuration   int    `json:"teleport_duration"`
	Coleaders          int    `json:"coleader_count"`
	Managers           int    `json:"manager_count"`
	Officers           int    `json:"officer_count"`
	Total              int    `json:"total_count"`
	RadiusClaimLimit   int    `json:"claim_radius_limit"`
	FMapRadius         int    `json:"fmap_radius"`
	FactionPermissions struct {
		Recruit   []string `json:"recruit"`
		Assistant []string `json:"assistant"`
		Officer   []string `json:"officer"`
		Manager   []string `json:"manager"`
		Coleader  []string `json:"coleader"`
		Leader    []string `json:"leader"`
	} `json:"faction_permissions"`
	Channels struct {
		Global    string `json:"global"`
		Truces    string `json:"truces"`
		Allies    string `json:"allies"`
		Faction   string `json:"faction"`
		Moderator string `json:"moderator"`
	} `json:"channels"`
}

//go:embed faction.json
var defaultFc []byte

// Faction is a global instance of the FactionConfig
var Faction FactionConfig

var rankIDs = map[int]string{
	0: "recruit",
	1: "assistant",
	2: "officer",
	3: "manager",
	4: "coleader",
	5: "leader",
}

var rankNames = map[int]string{
	0: "Recruit",
	1: "Assistant",
	2: "Officer",
	3: "Manager",
	4: "Co-Leader",
	5: "Leader",
}

// RankHasPermission returns whether a Faction rank has certain permission
func RankHasPermission(rank string, permission string) bool {
	permissions := GetFactionConfig[map[string]any]("faction_permissions")[rank].([]any)

	for _, p := range permissions {
		if p == permission || p == "all" {
			return true
		}

		if RankFromID(p.(string)) != -1 && RankHasPermission(p.(string), permission) {
			return true
		}
	}

	return false
}

// RankWithNativePermission returns the rankID of that rank which natively possesses a permission
func RankWithNativePermission(permission string) string {
	for _, r := range rankIDs {
		permissions := GetFactionConfig[map[string]any]("faction_permissions")[r].([]any)

		for _, p := range permissions {
			if p == permission {
				return r
			}
		}
	}

	return ""
}

// RankID returns the rank id
func RankID(rank int) string {
	return rankIDs[rank]
}

// RankName returns the rank name
func RankName(rank int) string {
	return rankNames[rank]
}

// RankFromID returns the rank from the id
func RankFromID(rankID string) int {
	for rank, r := range rankIDs {
		if r == rankID {
			return rank
		}
	}

	return -1
}

// ChannelFromID returns the channel from the id
func ChannelFromID(channelID string) chat.Channel {
	switch channelID {
	case "truces", "t":
		return chat.TrucesChannel{}
	case "allies", "a":
		return chat.AlliesChannel{}
	case "faction", "f":
		return chat.FactionChannel{}
	case "moderator", "m":
		return chat.ModeratorChannel{}
	}

	return chat.GlobalChannel{}
}
