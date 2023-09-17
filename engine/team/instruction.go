package team

import "encoding/json"

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
