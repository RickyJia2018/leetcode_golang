/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	left := 0
	right := len(nums)
	for left < right{
		mid:= left + (right-left)/2
		if nums[mid]>target{
			//on left side
			right = mid
		}else if nums[mid]<target {
			left = mid + 1
		}else{ //found the value
			return mid
		}
	}
	return right
}
// @lc code=end

// time complexity O(log n)  space O(1)
// unit test benchmarks