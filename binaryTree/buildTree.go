package binaryTree
/*
 * 前序遍历，中序遍历构造二叉树
 * leetcode: 105题
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val:rootVal}
	index := 0
	for i:=0; i<len(inorder); i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	maxLength := len(inorder) - 1
	if index != 0 {
		root.Left = buildTree(preorder[1: index + 1], inorder[:index])
	}
	if maxLength >= index + 1 {
		root.Right = buildTree(preorder[index + 1:], inorder[index + 1:])
	}
	return root
}