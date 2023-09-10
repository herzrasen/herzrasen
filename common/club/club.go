package club

import "github.com/herzrasen/common/player"

type Club struct {
	Name  string
	Squad []*player.Player
	// todo stadium, ...
}
