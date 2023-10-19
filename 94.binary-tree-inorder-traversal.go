/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
    var traversal func(node *TreeNode)
	traversal = func(node *TreeNode){
		if node == nil {
			return
		}else{
			traversal(node.Left)
			res = append(res, node.Val)
			traversal(node.Right)
		}
	}
	traversal(root)
	return res;
}
// @lc code=end

