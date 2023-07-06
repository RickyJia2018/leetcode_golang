//Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	left := 0
	right := len(nums) - 1
	index := len(nums) - 1

	result := make([]int, len(nums))

	for left < right {
		leftNum := nums[left] * nums[left]
		rightNum := nums[right] * nums[right]

		if leftNum < rightNum {
			result[index] = rightNum
			right--
			index--
		} else {
			result[index] = leftNum
			index--
			left++
		}
	}
	return result
}
func main() {
	nums := []int{-1, -4, 10, -9, 2, 4, 6}
	newNums := sortedSquares(nums)
	fmt.Println(newNums)
}
