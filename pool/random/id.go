package random

import "github.com/rs/xid"

func (r *Random) Id() string {
	return xid.New().String()
}
