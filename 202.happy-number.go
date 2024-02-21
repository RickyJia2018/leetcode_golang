/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	mySet := make(map[int]bool)

	for n != 1 {
		n = getSum(n)
		fmt.Println(n)
		if _, has := mySet[n]; has {
			break
		}
		mySet[n] = true
	}

	return n == 1

}
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// @lc code=end

