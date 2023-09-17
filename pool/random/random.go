package random

import (
	"github.com/herzrasen/common/gender"
	"github.com/herzrasen/pool/config"
)

type Random struct {
	minFloat   float32
	maxFloat   float32
	minInt     uint32
	maxInt     uint32
	Gender     gender.Gender
	NameConfig *config.NameConfig
	MinAttr    float32
	MaxAttr    float32
	MinYear    uint32
	MaxYear    uint32
	MinHeight  uint32
	MaxHeight  uint32
	MinWeight  uint32
	MaxWeight  uint32
}
