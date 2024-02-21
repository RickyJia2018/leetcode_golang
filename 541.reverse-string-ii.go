/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	res := []byte(s)

	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s) {
			myReverse(res[i : i+k])
		} else {
			myReverse(res[i:])
		}
	}
	return string(res)

}

func myReverse(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// @lc code=end

