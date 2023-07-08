package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	leftIndex := findLeftBound(nums, target)
	rightIndex := findRightBound(nums, target)
	return []int{leftIndex, rightIndex}
}

func findLeftBound(nums []int, target int) int {

	left := 0
	right := len(nums)
	result := -1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			result = mid
			right = mid

		}
	}

	return result
}

func findRightBound(nums []int, target int) int {
	left := 0
	right := len(nums)
	result := -1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			result = mid
			left = mid + 1
		}
	}
	return result

}
