package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	groupPrevious := dummy

	for true {
		kth := getKth(groupPrevious, k)

		if kth == nil {
			break
		}
		nodeAfterGroup := kth.Next
		previous, curr := kth.Next, groupPrevious.Next

		for curr != nodeAfterGroup {
			next := curr.Next

			curr.Next = previous
			previous = curr
			curr = next
		}
		reversedLastNode := groupPrevious.Next
		//This update the first node of the group next pointer to the kth of the next one
		groupPrevious.Next = kth
		groupPrevious = reversedLastNode
	}
	return dummy.Next
}

func getKth(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}
