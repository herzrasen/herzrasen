package random

func (r *Random) Attr() float32 {
	attrRandom := Random{minFloat: r.MinAttr, maxFloat: r.MaxAttr}
	return attrRandom.Float()
}
