/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
    res := len(nums) + 1
	left, right := 0,0
	sum := 0
	for ;right< len(nums); right++ {
		sum += nums[right]
		
		for sum>=target {
			subLen := right - left + 1
			if subLen < res {
				res = subLen
			}

			sum -= nums[left]
			left ++

		}
		
	}
	if res == len(nums) + 1 {
		res = 0
	}
	return res
}
// @lc code=end

