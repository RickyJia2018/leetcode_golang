/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	myMap := make(map[int]int)
	count := 0
	// 遍历nums1和nums2数组，统计两个数组元素之和，和出现的次数，放到map中

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			myMap[num1+num2]++
		}
	}
	// 遍历nums3和nums4数组，找到如果 0-(c+d) 在map中出现过的话，就把map中key对应的value也就是出现次数统计出来,

	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			count += myMap[-num3-num4]
		}
	}

	return count
}

// @lc code=end

