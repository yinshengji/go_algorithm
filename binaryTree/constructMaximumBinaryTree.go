package binaryTree

/*
 * 构造最大二叉树
 * leetcode: 654题
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal := -1 << 63
	index := 0
	for i := 0; i < len(nums) ; i++  {
		if nums[i] >= maxVal {
			maxVal = nums[i]
			index = i
		}
	}
	root := &TreeNode{Val:maxVal}
	root.Left = constructMaximumBinaryTree(nums[0:index])
	if index + 1 > len(nums) {
		root.Right = nil
	} else {
		root.Right = constructMaximumBinaryTree(nums[index+1:])
	}

	return root
}
