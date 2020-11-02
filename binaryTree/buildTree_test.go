package binaryTree

import (
	"github.com/kr/pretty"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{1, 2, 3}
	inorder := []int{2, 3, 1}
	root := buildTree(preorder, inorder)
	pretty.Println(root)
}