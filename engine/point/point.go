package point

import "math"

type Point struct {
	X int8 `json:"x"`
	Y int8 `json:"y"`
}

func (p *Point) Distance(other *Point) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
