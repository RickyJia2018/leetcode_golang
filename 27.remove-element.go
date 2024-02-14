/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
    slow:=0
	fast:=0

	for fast<len(nums){
		if nums[fast] == val{
			fast ++
		}else{
			nums[slow] = nums[fast]
			fast++
			slow++
		}
	}
	return slow
}
// @lc code=end

/* all have fast++ -> move out */

// for fast<len(nums){
// 	if nums[fast]!=val{
// 		nums[slow]=nums[fast]
// 		slow++
// 	}
// 	fast++
// }
