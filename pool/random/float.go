package random

import "math/rand"

func (r *Random) Float() float32 {
	return r.minFloat + rand.Float32()*(r.maxFloat-r.minFloat)
}
