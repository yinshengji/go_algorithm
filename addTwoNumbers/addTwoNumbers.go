package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var currentNode *ListNode
	var tailNode *ListNode
	var head *ListNode
	head = new(ListNode)
	if l1.Val == 0 && l1.Next == nil {
		head = l2
		return head
	}
	if l2.Val == 0 && l2.Next == nil {
		head = l1
		return head
	}

	var carry = 0
	currentL1Node := l1
	currentL2Node := l2
	currentNode = head
	tailNode = head
	for currentL1Node != nil {
		currentNode = new(ListNode)
		currentNode.Val = currentL1Node.Val + carry
		if currentL2Node != nil {
			currentNode.Val += currentL2Node.Val
		}
		carry = 0
		if currentNode.Val >= 10 {
			carry = 1
			currentNode.Val = currentNode.Val % 10
		}
		currentL1Node = currentL1Node.Next
		if currentL2Node != nil {
			currentL2Node = currentL2Node.Next
		}
		tailNode.Next = currentNode
		tailNode = tailNode.Next
	}

	for currentL2Node != nil {
		currentNode = new(ListNode)
		currentNode.Val = currentL2Node.Val + carry
		carry = 0
		if currentNode.Val >= 10 {
			carry = 1
			currentNode.Val = currentNode.Val % 10
		}
		currentL2Node = currentL2Node.Next
		tailNode.Next = currentNode
		tailNode = tailNode.Next
	}

	if carry >= 1 {
		currentNode = &ListNode{Val:carry, Next:nil}
		tailNode.Next = currentNode
	}
	return head.Next
}