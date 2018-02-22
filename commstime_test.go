package commstime

import "testing"

func BenchmarkCommstime(bench *testing.B) {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	d := make(chan int)

	go prefix(c, a, 0)
	go delta(a, b, d)
	go succ(b, c)

	consume(d, bench.N)
}
