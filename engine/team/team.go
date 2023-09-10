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
	Player  *player.Player
	Warning bool
	Card    *Card
}

type Substitution struct {
	SubbedOffPlayerId string
	BroughtOnPlayerId string
	Second            uint32
}

type Team struct {
	ActivePlayers []ActivePlayer   `json:"activePlayers"`
	Squad         []*player.Player `json:"squad"`
	Substitutions []Substitution   `json:"substitutions"`
	Bookings      []Booking        `json:"bookings"`
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
