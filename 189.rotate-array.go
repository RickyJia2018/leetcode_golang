/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int)  {
	if len(nums) <2 {
		return
	}
	k=k%len(nums)
	reverse(nums,0,len(nums) - k-1)
	reverse(nums,len(nums) - k,len(nums)-1)
	reverse(nums,0,len(nums)-1)
}

func reverse(nums []int, start int, end int)  {
	left,right := start,end
	for left<right{
		nums[left],nums[right] = nums[right],nums[left]
        left++
        right--
	}

}
// @lc code=end

