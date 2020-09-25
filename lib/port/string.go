package port

import "strconv"

func (p Port) String() string {
	return strconv.Itoa(int(p))
}