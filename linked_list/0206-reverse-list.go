package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var previous *ListNode

	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}

	return previous
}
