package commstime

func prefix(cin, cout chan int, v int) {
	cout <- v
	for v := range cin {
		cout <- v
	}
}

func id(cin, cout chan int) {
	for v := range cin {
		cout <- v
	}
}

func delta(cin, cout1, cout2 chan int) {
	for v := range cin {
		//go func(v int) { cout1 <- v }(v)
		cout1 <- v
		cout2 <- v
	}
}

func succ(cin, cout chan int) {
	for v := range cin {
		cout <- v + 1
	}
}

func consume(cin chan int, n int) {
	for i := 0; i < n; i++ {
		<-cin
	}
}
