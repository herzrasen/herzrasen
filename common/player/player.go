package player

import (
	"github.com/herzrasen/common/gender"
	"time"
)

type Player struct {
	Id             string         `json:"id"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Gender         gender.Gender  `json:"gender"`
	DateOfBirth    time.Time      `json:"dateOfBirth"`
	TechnicalAttrs TechnicalAttrs `json:"technicalAttrs"`
	PhysicalAttrs  PhysicalAttrs  `json:"physicalAttrs"`
	MentalAttrs    MentalAttrs    `json:"mentalAttrs"`
}

type TechnicalAttrs struct {
	Catching  float32 `json:"catching"`
	Dribbling float32 `json:"dribbling"`
	Heading   float32 `json:"heading"`
	Reflexes  float32 `json:"reflexes"`
	Shooting  float32 `json:"shooting"`
	Throwing  float32 `json:"throwing"`
	Passing   float32 `json:"passing"`
	Tackling  float32 `json:"tackling"`
	Technique float32 `json:"technique"`
}

type PhysicalAttrs struct {
	Acceleration    float32 `json:"acceleration"`
	Agility         float32 `json:"agility"`
	Balance         float32 `json:"balance"`
	Jumping         float32 `json:"jumping"`
	InjuryProneness float32 `json:"injuryProneness"`
	Pace            float32 `json:"pace"`
	Stamina         float32 `json:"stamina"`
	Strength        float32 `json:"strength"`
}

type MentalAttrs struct {
	Aggression    float32 `json:"aggression"`
	Anticipation  float32 `json:"anticipation"`
	Communication float32 `json:"communication"`
	Composure     float32 `json:"composure"`
	Decisions     float32 `json:"decisions"`
	Eccentricity  float32 `json:"eccentricity"`
	Focus         float32 `json:"focus"`
	Leadership    float32 `json:"leadership"`
	Positioning   float32 `json:"positioning"`
}
