package main

import (
	"flag"
	"fmt"
)

func main() {
	limit := flag.Int("limit", 1000, "upper limit for the calculation")
	flag.Parse()

	total := 0
	for i := 0; i < *limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	fmt.Println(total)
}
