package binaryTree

import (
	"github.com/kr/pretty"
	"testing"
)

func TestInvertTree(t *testing.T) {
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
	pretty.Println(root)
	newRoot := invertTree(root)
	pretty.Println(newRoot)
	//result := []int{10, 9, 8, 7, 6, 4, 3}
	if newRoot.Left.Val != 9 {
		t.Fail()
	}
	if newRoot.Right.Val != 8 {
		t.Fail()
	}
}
