/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	if len(s)!= len(goal) {
		return false
		
	}
	return strings.Contains(s+s, goal)
}
// @lc code=end

