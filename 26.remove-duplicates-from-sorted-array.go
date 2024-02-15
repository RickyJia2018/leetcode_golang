/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums)<1{
		return 0
	}
    slow:=0
	fast:=0

	for fast<len(nums){
		if nums[slow] == nums[fast]{
			nums[slow] = nums[fast]
		}else{
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow+1
}
// @lc code=end

