package linkedList
/**
 * leetcode: 206. 反转链表
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	var t *ListNode
	pre := head
	for pre != nil {
		t = pre.Next
		pre.Next = cur
		cur = pre
		pre = t
	}
	return cur
}