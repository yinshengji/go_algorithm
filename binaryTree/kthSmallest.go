package binaryTree

var (
	resultNum int
	indexTh int
)

func kthSmallest(root *TreeNode, k int) int {
	indexTh = 0
	resultNum = 0
	traverseTree(root, k)
	return resultNum
}

func traverseTree(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverseTree(root.Left, k)
	indexTh++
	if indexTh == k {
		resultNum = root.Val
		return
	}
	traverseTree(root.Right, k)
}
