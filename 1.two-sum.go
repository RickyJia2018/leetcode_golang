/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	result :=make([]int, 2)
	myHash := make(map[int]int)

	for i:=0; i<len(nums);i++{
		temp:= target - nums[i]

		index, ok := myHash[temp]
		if ok{
			result[0] = index
			result[1] = i
			return result
		}else{
			myHash[nums[i]] = i
		}
	}

	return result;
}
// @lc code=end

