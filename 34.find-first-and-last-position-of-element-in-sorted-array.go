/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
    left:=findLeft(nums, target)
	right:=findRight(nums, target)
	return []int{left,right}
}

func findLeft(nums []int, target int) int{
	if len(nums)<1{
		return -1
	}
	left:=0
	right:=len(nums)
	res:=-1
	
	for left<right{
		mid:=left + (right-left)/2
		if nums[mid]<target{
			left = mid + 1
		}else if nums[mid]>target{
			right = mid
		}else{
			//update res
			res = mid
			//keep moving
			right = mid 
		}
	}
	return res
}


func findRight(nums []int,target int) int {
	if len(nums)<1{
		return -1
	}
	left:=0
	right:=len(nums)
	res:=-1
	for left<right{
		mid:=left+(right-left)/2
		if nums[mid]>target{
			right = mid
		}else if nums[mid]<target{
			left = mid + 1
		}else{
			res = mid
			left = mid + 1
		}
	}
	return res
 }


// @lc code=end

