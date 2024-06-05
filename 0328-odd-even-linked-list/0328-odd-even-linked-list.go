func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nechetHead := &ListNode{Next: head.Next}
	chet := head
	nechet := head.Next

	for nechet != nil && nechet.Next != nil {
		chet.Next = nechet.Next
		chet = chet.Next
		nechet.Next = chet.Next
		nechet = nechet.Next
	}

	chet.Next = nechetHead.Next

	return head
}