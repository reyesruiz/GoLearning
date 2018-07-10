package lengthconv

import (
	"fmt"
)

type Feet float64
type Meters float64

func (meeter Meters) String() string {
	return fmt.Sprintf("%g Meeters", meeter)
}

func (feet Feet) String() string {
	return fmt.Sprintf("%g Feet", feet)
}
