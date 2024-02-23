/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {

	res := make([]int, 0)
	for i, j := 0, 0; i+j < n+m; {
		// check m,n first
		if i == m {
			res = append(res, nums2[j:]...)
			break
		}
		if j == n {
			res = append(res, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}

	}
	fmt.Println(res)
	copy(nums1, res) // nums1 = res won't affect outside nums1
}

// @lc code=end

