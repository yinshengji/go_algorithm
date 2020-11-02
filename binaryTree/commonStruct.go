package binaryTree

type TreeNode struct {
	Val  int
	Left *TreeNode
	Right * TreeNode
}

func (t *TreeNode) Print(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	queue := []*TreeNode{node}
	result := []int{}
	for len(queue) > 0 {
		if queue[0] != nil {
			result = append(result, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
		}
		queue = queue[1:]
	}
	return result
}
