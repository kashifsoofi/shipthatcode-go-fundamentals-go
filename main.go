package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(n)
	}
}
