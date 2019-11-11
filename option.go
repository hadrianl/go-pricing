package pricing

import (
	"fmt"

	"github.com/hadrianl/go-pricing/base/bs"
)

func Calc() {
	// price := bs.CalcPrice(100, 95, 0.02, 3/12, 0.5, 1)
	fmt.Println(bs.Measure(100, 95, 0.1, 3/12, 0.5, 1))
}
