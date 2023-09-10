package random

import "math/rand"

func (r *Random) FloatInRange(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
