package binaryTree

import (
	"github.com/kr/pretty"
	"testing"
)

/*
 * 前序遍历，中序遍历构造二叉树
 */

func TestBuildTreeByInorderAndPostOrder(t *testing.T) {
	inorder := []int{2, 3, 1, 4, 5}
	postorder := []int{2, 3, 4, 5, 1}
	root := buildTreeByInorderAndPostOrder(inorder, postorder)
	pretty.Println(root)
}