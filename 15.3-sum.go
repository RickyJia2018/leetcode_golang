/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int;
	for i:=0;i<len(nums);i++ {
		left:=i+1
		right:=len(nums) - 1
		sum :=0
		for left<right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
                left++
			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}
		}
	}
	return result

}
// @lc code=end

