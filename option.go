package option

import (
	"fmt"

	blackscholes "github.com/hadrianl/go-pricing/base/black-scholes"
)

func Calc() {
	// price := bs.CalcPrice(100, 95, 0.02, 3/12, 0.5, 1)
	fmt.Println(blackscholes.Measure(90, 100, 0.1, 1.0/12, 0.35, 0.05, 1))
}
