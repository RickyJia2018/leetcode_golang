/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start

/*  Is this a public api or for internal use? 
 *  need more speed or less space? Do I need consider Multi thread situation?
 *  Max & min length of nums? target always inside the range of numbers?
 *  What should I return if nums if null or length < 1?
 */

func search(nums []int, target int) int {
	if len(nums)<1{
		return -1
	}

	left:=0
	right:=len(nums)

	for left<right{
		mid:= left + (right-left)/2
		if nums[mid] > target{
			right = mid
		}else if nums[mid] < target{
			left = mid +1
		}else{
			return mid
		}
	}
	return -1
}

/*
 * I Could write unit test, add function comments, do benchmarking 
 */
// @lc code=end
