/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var res *ListNode
    sizeA := countNode(headA)
	sizeB := countNode(headB)
	if sizeA<0||sizeB<0{
		return res
	}
	var cur1 *ListNode
	var cur2 *ListNode
	if sizeA > sizeB{
		cur1 = headA
		cur2 = headB
	}else{
		cur1 = headB
		cur2 = headA
	}

	diff := int(math.Abs(float64(sizeA-sizeB)))
	for i:=0;i<diff;i++{
		cur1 = cur1.Next
	}

	for cur1 !=nil{
		if cur1 == cur2{
			res = cur1
			break
		}else{
			cur1=cur1.Next
			cur2=cur2.Next
		}
	}
	return res

}

func countNode(head *ListNode) int{
	if head == nil{
		return -1
	}
	count:=0
	for head!=nil{
		count++
		head = head.Next
	}
	return count
}
// @lc code=end

