package main

import "fmt"

func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
func main() {

	nums := []int{1, 3, 5, 6}
	index := searchInsert(nums, 5)
	fmt.Println(index)

}
