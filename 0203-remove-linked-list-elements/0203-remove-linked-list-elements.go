func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
		return head
	}
	freePointer := &ListNode{Next: head}
	endPointer := freePointer
	for endPointer.Next != nil {
		if endPointer.Next.Val == val {
			endPointer.Next = endPointer.Next.Next
		} else {
			endPointer = endPointer.Next
		}
	}
	return freePointer.Next
}