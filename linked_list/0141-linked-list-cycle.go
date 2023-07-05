package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	tortoise, hare := head, head

	for hare != nil && hare.Next != nil {
		hare = hare.Next.Next
		tortoise = tortoise.Next

		if hare == tortoise {
			return true
		}
	}

	return false
}
