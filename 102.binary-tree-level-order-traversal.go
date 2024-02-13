/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return result
	}
	result := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	size := 1
	queue = append(queue, root)
	for len(queue) > 0 {
		count := 0
		row := make([]int, 0)
		for size > 0 {
			tempNode := queue[0]
			queue = queue[1:]
			row = append(row, tempNode.Val)
			size--
			if tempNode.Left != nil {
				queue = append(queue, tempNode.Left)
				count++
			}
			if tempNode.Right != nil {
				queue = append(queue, tempNode.Right)
				count++
			}

		}
		size = count
		count = 0
		result = append(result, row)
	}
	return result
}

// @lc code=end

