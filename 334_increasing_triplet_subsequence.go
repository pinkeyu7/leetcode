package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	a := math.MaxInt64
	b := math.MaxInt64

	for _, num := range nums {
		if num <= a {
			a = num
		} else if num <= b {
			b = num
		} else {
			return true
		}
		fmt.Println("a:", a, "b:", b)
	}

	return false
}
