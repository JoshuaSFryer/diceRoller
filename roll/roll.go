// roll represents a single die roll. It contains the value of the roll, and
// a flag indicating whether that roll was a critical success (1), critical
// failure(-1), or neither.
package roll

import (
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"errors"
)

type Roll struct {
	Value    int
	Critical int
}

func New(val int, crit int) (Roll, error) {
	r := Roll{val, crit}
	return r, nil
}

func (r Roll) Print() {
	if r.Critical == 1 {
		// Critical successes are highlighted green
		ct.Foreground(ct.Green, true)
	} else if r.Critical == -1 {
		// Critical failures are highlighted red
		ct.Foreground(ct.Red, true)
	} else {
		// No highlight for regular rolls
		ct.ResetColor()
	}
	fmt.Printf("%d ", r.Value)
	ct.ResetColor()
}
