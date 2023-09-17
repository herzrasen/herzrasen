package team

import (
	"encoding/json"
	"strings"
)

type BasePosition uint8

const (
	GoalKeeper BasePosition = 0
	Defender   BasePosition = 1
	Midfielder BasePosition = 2
	Striker    BasePosition = 3
	Unknown    BasePosition = 4
)

func (b *BasePosition) MarshalJSON() ([]byte, error) {
	var s string
	switch *b {
	case GoalKeeper:
		s = "goalkeeper"
	case Defender:
		s = "defender"
	case Midfielder:
		s = "midfielder"
	case Striker:
		s = "striker"
	default:
		s = "unknown"
	}
	return json.Marshal(s)
}

func (b *BasePosition) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "goalkeeper":
		*b = GoalKeeper
	case "defender":
		*b = Defender
	case "midfielder":
		*b = Midfielder
	case "striker":
		*b = Striker
	default:
		*b = Unknown
	}
	return nil
}
