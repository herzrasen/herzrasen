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

type Instruction uint8

const (
	MaintainPosition Instruction = 0
	PassTheBall      Instruction = 1
	TryToFinish      Instruction = 2
	PlayForward      Instruction = 3
	PlayCarefully    Instruction = 4
	PushUp           Instruction = 5
	AttackEarly      Instruction = 6
)

type Position struct {
	BasePosition BasePosition  `json:"basePosition"`
	Instructions []Instruction `json:"instructions"`
}

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

func (i *Instruction) MarshalJSON() ([]byte, error) {
	var s string
	switch *i {
	case MaintainPosition:
		s = "maintainPosition"
	case PassTheBall:
		s = "passTheBall"
	case TryToFinish:
		s = "tryToFinish"
	case PlayForward:
		s = "playForward"
	case PlayCarefully:
		s = "playCarefully"
	case PushUp:
		s = "pushUp"
	case AttackEarly:
		s = "attackEarly"
	default:
		s = "unknown"
	}
	return json.Marshal(s)
}

func (i *Instruction) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "maintainPosition":
		*i = MaintainPosition
	case "passTheBall":
		*i = PassTheBall
	case "tryToFinish":
		*i = TryToFinish
	case "playForward":
		*i = PlayForward
	case "playCarefully":
		*i = PlayCarefully
	case "pushUp":
		*i = PushUp
	case "attackEarly":
		*i = AttackEarly
	}
	return nil
}
