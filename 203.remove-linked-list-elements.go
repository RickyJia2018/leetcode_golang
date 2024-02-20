/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	dumHead:=&ListNode{Next:head}
	cur := dumHead

	for cur.Next !=nil{
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return dumHead.Next
}
// @lc code=end

