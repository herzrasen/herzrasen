package team

import (
	"encoding/json"
	"github.com/herzrasen/common/player"
	"os"
)

type ActivePlayer struct {
	Position Position `json:"position"`
	PlayerId string   `json:"playerId"`
}

type Card uint8

const (
	Yellow    Card = 0
	YellowRed Card = 1
	Red       Card = 2
)

type Booking struct {
	PlayerId string `json:"playerId"`
	Warning  bool   `json:"warning"`
	Card     *Card  `json:"card"`
}

type Substitution struct {
	SubbedOffPlayerId string `json:"subbedOffPlayerId"`
	BroughtOnPlayerId string `json:"broughtOnPlayerId"`
	Second            uint32 `json:"second"`
}

type Team struct {
	Id            string          `json:"id"`
	ActivePlayers []ActivePlayer  `json:"activePlayers"`
	Squad         []player.Player `json:"squad"`
	Substitutions []Substitution  `json:"substitutions"`
	Bookings      []Booking       `json:"bookings"`
	// substitutionRules, ...
}

func FromFile(teamFile string) (*Team, error) {
	data, err := os.ReadFile(teamFile)
	if err != nil {
		return nil, err
	}
	var team Team
	err = json.Unmarshal(data, &team)
	if err != nil {
		return nil, err
	}
	return &team, nil
}
