package addTwoNumbers

import (
	"github.com/kr/pretty"
	"testing"
)

func TestTest(t *testing.T)  {
	//1->3->6->1
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = &ListNode{Val:3, Next:nil}
	l1.Next.Next = &ListNode{Val:6, Next:nil}
	l1.Next.Next.Next = &ListNode{Val:1, Next:nil}
	//2->4->3
	l2 := new(ListNode)
	l2.Val = 2
	l2.Next = &ListNode{Val:4, Next:nil}
	l2.Next.Next = &ListNode{Val:3, Next:nil}
	l3 := AddTwoNumbers(l1, l2)
	pretty.Print(l3)
	//l1.Next.Next = &ListNode{Val:3, Next:nil}
	//pretty.Println(l3)
}
