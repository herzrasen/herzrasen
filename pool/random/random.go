package random

import (
	"github.com/herzrasen/common/gender"
	"github.com/herzrasen/pool/config"
)

type Random struct {
	Gender     gender.Gender
	NameConfig *config.NameConfig
}
