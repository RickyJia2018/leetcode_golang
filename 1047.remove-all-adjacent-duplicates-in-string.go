/*
 * @lc app=leetcode id=1047 lang=golang
 *
 * [1047] Remove All Adjacent Duplicates In String
 */

// @lc code=start
func removeDuplicates(s string) string {

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
	// stack := make([]byte, 0)
	// for i := 0; i < len(s); i++ {
	// 	if len(stack) > 0 && stack[len(stack)-1] == s[i] {
	// 		stack = stack[:len(stack)-1]
	// 	} else {
	// 		stack = append(stack, s[i])
	// 	}
	// }
	// return string(stack)
}

// @lc code=end

