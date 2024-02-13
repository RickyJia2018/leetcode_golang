/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
    left:=findLeft(nums, target)
	right:=findRight(nums, target)
	return []int{left,right}
}
func findLeft(nums []int,target int) int {
	if len(nums) == 0 { return -1 }
	left ,right := 0, len(nums)
	res := -1

	for left<right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		}else if nums[mid] < target{
			left = mid + 1
		}else{
			right = mid
			res = mid
		}

	}

	return res
}
func findRight(nums []int,target int) int {
	if len(nums) == 0 { return -1 }
	left ,right := 0, len(nums)
	res := -1
	for left<right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		}else if nums[mid] < target{
			left = mid + 1
		}else{
			left = mid + 1
			res = mid
		}
	}
	return res
}
// @lc code=end

