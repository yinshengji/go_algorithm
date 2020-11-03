package binaryTree

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	index map[string]int
	result []*TreeNode
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	index = make(map[string]int)
	result = nil
	traverse(root)
	return result
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := traverse(root.Left)
	right := traverse(root.Right)
	str := left + "," + right + "," + strconv.Itoa(root.Val)
	count, ok := index[str]
	if ok {
		if count == 1 {
			result = append(result, root)
		}
		index[str] += 1
	} else {
		index[str] = 1
	}

	return str
}



