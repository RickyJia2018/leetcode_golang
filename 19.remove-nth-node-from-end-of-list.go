/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // what to do if n > length of the list??

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
	if head == nil || n<=0{
		return head
	}
	dumHead := &ListNode{Next: head}
	temp := dumHead
	for i:=0;i<n; i++{
		temp = temp.Next
	}

	pre:=dumHead
	cur:=pre.Next

	for temp.Next!=nil{
		pre = cur
		cur = cur.Next
		temp = temp.Next
	}
	pre.Next = cur.Next
	return dumHead.Next
}
// @lc code=end

