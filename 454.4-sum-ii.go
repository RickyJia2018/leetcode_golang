/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	myMap:=make(map[int]int)
	count:=0
	for _,num1 := range nums1{
		for _,num2 := range nums2{
			myMap[num1+num2]++
		}
	}

	for _,num3 := range nums3{
		for _,num4 := range nums4{
			count +=myMap[-num3-num4]
		}
	}
	return count
}
// @lc code=end

