func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode
	cur := slow
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	for prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}
