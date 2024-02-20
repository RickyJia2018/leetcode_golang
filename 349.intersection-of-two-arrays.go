/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return

	// temp := make(map[int]int)
	// res := make([]int, 0)
	// for _, n := range nums1 {
	// 	for _, m := range nums2 {
	// 		if n == m {
	// 			temp[n]++
	// 		}
	// 	}
	// }

	// for key, _ := range temp {
	// 	res = append(res, key)
	// }

	// return res
}

// @lc code=end

