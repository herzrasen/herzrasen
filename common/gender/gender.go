package gender

import (
	"encoding/json"
	"strings"
)

type Gender uint8

const (
	Male      Gender = 0
	Female    Gender = 1
	Undefined Gender = 2
)

func FromString(g string) Gender {
	switch g {
	case "male":
		return Male
	case "female":
		return Female
	default:
		return Undefined
	}
}

func (g *Gender) MarshalJSON() ([]byte, error) {
	var s string
	switch *g {
	case Male:
		s = "male"
	case Female:
		s = "female"
	default:
		s = "undefined"
	}
	return json.Marshal(s)
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "male":
		*g = Male
	case "female":
		*g = Female
	default:
		*g = Undefined
	}
	return nil
}
