/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil{
		return head
	}
	//dummy head make it easier
	dumHead := &ListNode{
		Next: head,
	}

	cur:=dumHead  
	for cur.Next!=nil && cur.Next.Next!=nil{
		firstNode := cur.Next
		secondNode := cur.Next.Next
		nextNode := secondNode.Next
		
		// KEY LOGIC: draw how the order changed
		cur.Next = secondNode
		secondNode.Next = firstNode
		firstNode.Next = nextNode
		cur = firstNode
	}
	return dumHead.Next
}
// @lc code=end

