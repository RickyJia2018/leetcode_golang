package main

import (
	"fmt"
	"sort"
)

func BinarySearch(nums []int, target int) int {

	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}

	return -1
}
func main() {
	var arr []int = []int{1, 2, 3, 4, 5, 6}
	sort.Ints(arr)
	fmt.Println(arr)
	target := 4
	index := BinarySearch(arr, target)
	fmt.Println(index)
}
