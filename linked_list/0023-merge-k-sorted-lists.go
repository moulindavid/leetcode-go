package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var mergedLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			list1 := lists[i]
			var list2 *ListNode
			if (i + 1) < len(lists) {
				list2 = lists[i+1]
			}
			mergedLists = append(mergedLists, mergeList(list1, list2))
		}
		lists = mergedLists
	}
	return lists[0]
}

func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	pointer1, pointer2 := list1, list2

	result := new(ListNode)

	temp := result

	for pointer1 != nil && pointer2 != nil {
		if pointer1.Val <= pointer2.Val {
			temp.Next = pointer1
			temp = temp.Next
			pointer1 = pointer1.Next
		} else {
			temp.Next = pointer2
			temp = temp.Next
			pointer2 = pointer2.Next
		}
	}

	for pointer1 != nil {
		temp.Next = pointer1
		temp = temp.Next
		pointer1 = pointer1.Next
	}
	for pointer2 != nil {
		temp.Next = pointer2
		temp = temp.Next
		pointer2 = pointer2.Next
	}
	return result.Next
}
