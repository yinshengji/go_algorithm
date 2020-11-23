package binaryTree

import "testing"

func TestMaxPathSum(t *testing.T)  {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:-2,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	max := maxPathSum(root)
	if max != 12 {
		t.Fail()
	}
}