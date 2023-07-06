// Leetcode 27

package main

func removeElement(nums []int, val int) int {

	slow := 0
	fast := 0

	for fast < len(nums) {
		if nums[fast] == val {
			fast++
		} else {
			nums[slow] = nums[fast]
			fast++
			slow++
		}

	}
	return slow
}
func main() {
	nums := []int{1, 2, 2, 3, 1, 1, 4, 5}
	removeElement(nums, 1)
}
