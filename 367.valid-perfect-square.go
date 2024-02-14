/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
    if num < 0{
		// Throw err	
		return false 
	}

	left := 0
	right := num + 1
	for left < right{
		mid := left + (right - left) /2 
		temp := mid*mid
		if temp == num{
			return true
		}else if temp < num{
			left = mid + 1
		}else{
			right = mid
		}
	}
	return false
}
// @lc code=end

