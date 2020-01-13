package main

import "fmt"

func main() {
	x := 600851475143
	div := 2
	for x > 1 {
		if x%div == 0 {
			x /= div
			continue
		}
		div++
	}

	fmt.Println(div)
}
