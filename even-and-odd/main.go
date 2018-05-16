package main

import (
	"fmt"
)

func main() {
	r := newRangeToSlice(0, 10)

	for _, val := range r {
		if val%2 == 0 {
			fmt.Printf("%v is even\n", val)
		} else {
			fmt.Printf("%v is odd\n", val)
		}

	}
}

func newRangeToSlice(start int, end int) []int {
	if start > end {
		fmt.Println("Start value cannot be greater than end value")
	}
	s := []int{}

	for start <= end {
		s = append(s, start)
		start++
	}

	return s
}
