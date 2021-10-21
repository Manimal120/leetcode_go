package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummyHead.Next
}
