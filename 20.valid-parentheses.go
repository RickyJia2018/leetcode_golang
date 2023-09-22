/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
    if len(s) == 0 {
		return true
	}
	hashMap := map[byte]byte{
		'(': ')',
		'{': '}',
        '[': ']',
	}
	stack := make([]byte,0)


	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] =='['{
			stack = append(stack, hashMap[s[i]])
		}else{
			if len(stack) == 0 {
				return false
			}
			if s[i] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}
	
    }
	return len(stack) == 0

}
func isValid2(s string) bool {
    hash := map[byte]byte{')':'(', ']':'[', '}':'{'}
    stack := make([]byte, 0)
    if s == "" {
        return true
    }

    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack = append(stack, s[i])
        } else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
            stack = stack[:len(stack)-1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}
// @lc code=end

