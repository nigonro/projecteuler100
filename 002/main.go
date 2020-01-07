package main

import "fmt"

func fibonacci(ch chan uint64, quit chan struct{}) {
	x, y := uint64(1), uint64(2)
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

func main() {
	ch := make(chan uint64)
	quit := make(chan struct{})

	sum := uint64(0)

	go func() {
		for {
			term := <-ch
			if term >= 4000000 {
				break
			}
			if term%2 == 0 {
				sum += term
			}
		}
		quit <- struct{}{}
	}()
	fibonacci(ch, quit)
	fmt.Println(sum)
}
