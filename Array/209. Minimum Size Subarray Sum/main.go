package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	var sum int = 0
	left := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			subLen := right - left + 1
			if subLen < result {
				result = subLen
			}
			sum -= nums[left]
			left++
		}
	}

	if result == len(nums)+1 {
		result = 0
	}

	return result

}
func main() {
	target := 4
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	min := minSubArrayLen(target, nums)
	fmt.Println(min)
}
