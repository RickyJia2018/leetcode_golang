/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start

/*
 * Who gonna use this? 
 * x=0?
 * 
 */
func mySqrt(x int) int {
	if x <= 0{
		return 0
	}
	left := 1
	right := x + 1
    root := 1
	for left<right{
		mid := left + (right - left)/2
		val := mid * mid
		if val > x{
			right = mid
		}else if val < x{
            left = mid + 1
            if mid>root{
                root = mid
            }
		}else{
            return mid
        }
	}
	return root
}
// @lc code=end

