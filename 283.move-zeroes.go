/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
	if len(nums)<1{
		return
	}
    slow,fast:=0,0

	for fast< len(nums){
		if nums[fast] == 0{
			fast++
		}else{
			nums[slow] = nums[fast]
			fast++
			slow++
		}
	}
	for slow<len(nums){
		nums[slow] = 0
		slow++
	}
}
// @lc code=end

