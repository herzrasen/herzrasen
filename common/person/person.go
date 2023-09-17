package person

import (
	"github.com/herzrasen/common/gender"
	"time"
)

type Person struct {
	Id          string        `json:"id"`
	Firstname   string        `json:"firstName"`
	Lastname    string        `json:"lastName"`
	Gender      gender.Gender `json:"gender"`
	DateOfBirth time.Time     `json:"dateOfBirth"`
	Height      Height        `json:"height"`
	Weight      Weight        `json:"weight"`
}

type Height struct {
	Centimeters uint32 `json:"cm"`
}

type Weight struct {
	Grams uint32 `json:"gr"`
}
