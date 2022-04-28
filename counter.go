package packa

import (
	"github.com/mosceo/counter"
)

func Inc() {
	counter.Inc()
}

func Dec() {
	counter.Dec()
}

func Counter() int {
	return counter.Counter
}

func init() {
	counter.Inc()

	counter.AddValue(1)
	counter.AddValue(2)
	counter.AddValue(3)
}
