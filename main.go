package main

import "fmt"

func main() {
	var w, h int
	fmt.Scan(&w)
	fmt.Scan(&h)
	area := w * h
	fmt.Println(area)
}
