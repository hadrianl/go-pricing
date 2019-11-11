package option

import (
	"fmt"

	"github.com/hadrianl/go-pricing/base/bs"
)

func Calc() {
	// price := bs.CalcPrice(100, 95, 0.02, 3/12, 0.5, 1)
	fmt.Println(bs.Measure(90, 100, 0.1, 1.0/12, 0.35, 0.05, 1))
}
