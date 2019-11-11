package option_test

import (
	"fmt"
	"testing"

	"github.com/hadrianl/go-pricing/base/bs"
)

func TestCalc(t *testing.T) {

	fmt.Println(bs.Measure(90.0, 100, 0.1, 3.0/12, 0.35, 0.05, -1))
}

func BenchmarkCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs.Measure(90.0, 100, 0.1, 3.0/12, 0.35, 0.05, -1)
	}
}
