package main

import "fmt"

func mySqrt(x int) int {

	left := 0
	right := x + 1 // !important

	result := 0

	for left < right {
		mid := left + (right-left)/2
		fmt.Println(mid)
		midNum := mid * mid

		if midNum <= x {

			result = mid
			left = mid + 1
		} else {
			right = mid
		}
	}
	return result
}
