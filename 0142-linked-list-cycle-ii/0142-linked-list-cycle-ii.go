func detectCycle(head *ListNode) *ListNode {
    first := head
	second := head

	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
		if first == second {
			break
		}
	}
	if second == nil || second.Next == nil {
		return nil
	}
	first = head
	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}