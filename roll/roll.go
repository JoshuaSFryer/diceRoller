package roll

import (
	"fmt"
	"github.com/daviddengcn/go-colortext"
)

type Roll struct {
	Value    int
	Critical int
}

func New(val int, crit int) Roll {
	r := Roll{val, crit}
	return r
}

func (r Roll) Print() {
	if r.Critical == 1 {
		ct.Foreground(ct.Green, true)
	} else if r.Critical == -1 {
		ct.Foreground(ct.Red, true)
	} else {
		ct.ResetColor()
	}
	fmt.Printf("%d ", r.Value)
    ct.ResetColor()
}
