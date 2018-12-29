// roll represents a single die roll. It contains the value of the roll, and
// a flag indicating whether that roll was a critical success (1), critical
// failure(-1), or neither.
package roll

import (
	"fmt"
	"github.com/daviddengcn/go-colortext"
)

// Ugly code duplication, but I don't know a better way to share constants
// between packages when it's something so small.
const (
	CritFailure int = iota - 1
	CritNeutral
	CritSuccess
)

type Roll struct {
	Value    int
	Critical int
}

func New(val int, crit int) (Roll) {
	r := Roll{val, crit}
	return r
}

func (r Roll) Print() {
	if r.Critical ==  CritSuccess{
		// Critical successes are highlighted green
		ct.Foreground(ct.Green, true)
	} else if r.Critical == CritFailure {
		// Critical failures are highlighted red
		ct.Foreground(ct.Red, true)
	} else {
		// No highlight for regular rolls
		ct.ResetColor()
	}
	fmt.Printf("%d ", r.Value)
	ct.ResetColor()
}
