/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	outer:
		for i := 0; i+m <= n; i++ {
			for j := range needle {
				if haystack[i+j] != needle[j] {
					continue outer
				}
			}
			return i
		}
		return -1
}
// @lc code=end

