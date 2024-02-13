/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	
}
// @lc code=end

// if len(nums) == 0 {
// 	return -1
// }
// left := 0
// right := len(nums)
// for left < right {
// 	mid:= left + (right-left)/2

// 	if nums[mid] < target {
// 		left = mid +1
// 	}else if nums[mid] > target {
// 		right = mid
// 	}else{
// 		return mid
// 	}
// }
// return right