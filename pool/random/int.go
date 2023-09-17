package random

func (r *Random) Int() int {
	floatRandom := Random{minFloat: float32(r.minInt), maxFloat: float32(r.maxInt)}
	return int(floatRandom.Float())
}
