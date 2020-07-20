package main

import "github.com/kr/pretty"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var currentNode *ListNode
	var tailNode *ListNode
	head := new(ListNode)
	head.Val = 1
	tailNode = head

	currentNode = new(ListNode)
	currentNode.Val = 2
	currentNode.Next = new(ListNode)
	tailNode.Next = currentNode
	tailNode = tailNode.Next

	currentNode = new(ListNode)
	currentNode.Val = 3
	currentNode.Next = new(ListNode)
	tailNode.Next = currentNode
	tailNode = tailNode.Next

	pretty.Println(currentNode, tailNode, head)
}
