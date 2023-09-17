package random

import "github.com/herzrasen/common/person"

func (r *Random) Person() person.Person {
	return person.Person{
		Id:          r.Id(),
		Firstname:   r.Firstname(),
		Lastname:    r.Lastname(),
		Gender:      r.Gender,
		DateOfBirth: r.Date(),
		Height:      r.Height(),
		Weight:      r.Weight(),
	}
}
