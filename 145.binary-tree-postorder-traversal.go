/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res:= make([]int,0)
	if root == nil{
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len()>0{
		lastNode:= stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		// res = append(res, lastNode.Val) // need reverse final res
		res = append([]int{lastNode.Val},res...)
		if lastNode.Left!=nil{
			stack.PushBack(lastNode.Left)
		}
		if lastNode.Right!=nil{
			stack.PushBack(lastNode.Right)
		}		

	}
	// reverse(res)
	// traversal(root, &res)
	return res
}
func reverse(arr []int){
	left := 0
	right := len(arr) -1
	for left < right{
		arr[left],arr[right] = arr[right], arr[left]
		left ++
		right --
	}
}
func traversal(node *TreeNode, res *[]int){
	if node == nil{
		return 
	}
	traversal(node.Left,res)
	traversal(node.Right,res)
	*res = append(*res,node.Val)
}
// @lc code=end

