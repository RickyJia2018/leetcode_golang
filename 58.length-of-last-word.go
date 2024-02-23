/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			left = i
			for i < len(s) && s[i] != ' ' {
				i++
			}
			right = i
		}
	}
	return right - left
}

// @lc code=end

