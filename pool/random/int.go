package random

func (r *Random) IntInRange(min int, max int) int {
	return int(r.FloatInRange(float32(min), float32(max)))
}
