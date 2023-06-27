package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	middle := findMiddle(head)
	reversed := reverse(middle.Next)

	// we dont want the head to continue after the middle value.
	middle.Next = nil
	merge(head, reversed)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(node *ListNode) *ListNode {
	var prev, curr *ListNode = nil, node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev

}

func merge(curr *ListNode, reversed *ListNode) {
	for curr != nil && reversed != nil {
		currNext := curr.Next
		reversedNext := reversed.Next

		curr.Next = reversed
		reversed.Next = currNext

		curr = currNext
		reversed = reversedNext
	}
}
