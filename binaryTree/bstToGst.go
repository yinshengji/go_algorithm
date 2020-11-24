package binaryTree

/*
 * leetcode 1038. 把二叉搜索树转换为累加树
 * 解题思路：中序遍历有序，累加树计算是倒叙，所以用右中左遍历方法
 */
var (
	total int = 0
)

func bstToGst(root *TreeNode) *TreeNode {
	total = 0
	traverseBST(root)
	return root
}

func traverseBST(root *TreeNode) {
	if root == nil {
		return
	}

	traverseBST(root.Right)
	total += root.Val
	root.Val = total
	traverseBST(root.Left)
}
