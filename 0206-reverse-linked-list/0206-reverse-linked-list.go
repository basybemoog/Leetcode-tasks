
func reverseList(head *ListNode) *ListNode {
     var reverse *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = reverse
		reverse = current
		current = next
	}
	return reverse  

}