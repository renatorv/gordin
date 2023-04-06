package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
