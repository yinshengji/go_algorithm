package binaryTree

import (
	"testing"
)

func TestTreeNode_Print(t *testing.T) {
	root := new(TreeNode)
	root.Val = 10
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.Val = 8
	root.Left.Left = new(TreeNode)
	root.Left.Right = new(TreeNode)
	root.Right.Val = 9
	root.Right.Left = new(TreeNode)
	root.Right.Right = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Right.Val = 4
	root.Right.Left.Val = 6
	root.Right.Right.Val = 7
	root.Print(root)
	//pretty.Println(nums)
}