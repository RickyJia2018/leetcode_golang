/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
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
func invertTree(root *TreeNode) *TreeNode {
	root = invert(root)
	return root
}

func invert(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	tempNode := node.Left
	node.Left = invert(node.Right)
	node.Right = invert(tempNode)
	return node
}

// @lc code=end

