package binaryTree

/*
 * 前序遍历，中序遍历构造二叉树
 * leetcode: 106题
 */

func buildTreeByInorderAndPostOrder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	index := 0
	rootVal := postorder[len(postorder) - 1]
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			index = i
			break;
		}
	}
	root := &TreeNode{Val:rootVal}
	maxIndex := len(inorder) - 1
	if index != 0 {
		root.Left = buildTreeByInorderAndPostOrder(inorder[0:index], postorder[0:index])
	}
	if maxIndex >= index + 1 {
		root.Right = buildTreeByInorderAndPostOrder(inorder[index + 1:], postorder[index:len(postorder) - 1])
	}

	return root
}