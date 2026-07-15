package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	total := 0
	for i := 0; i <= n; i++ {
		total += i
	}

	fmt.Println(total)
}
