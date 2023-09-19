/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x<0 {return false}
	temp := x
	newNum :=0

	for temp>0{
		newNum = newNum*10 + temp%10
		temp = temp/10
	}
	if x == newNum{
		return true
	}else{
		return false
	}
}
// @lc code=end

