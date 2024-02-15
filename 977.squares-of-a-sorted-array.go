/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
    left:=0
	right:=len(nums)-1
	index := right
	res := make([]int,len(nums))  // need a temp array to store result

	for index>=0{  //don't forget index 0
	
		leftNum :=nums[left]*nums[left]
		rightNum :=nums[right]*nums[right]

		if leftNum<=rightNum{
			res[index] = rightNum
			right--
		}else {
			res[index] = leftNum
			left++
		}
		index --
		fmt.Println(index)
		
	}
	return res
}
// @lc code=end

