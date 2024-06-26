func hasCycle(head *ListNode) bool {
	first := head
  second := head
  for second != nil && second.Next != nil{
    first = first.Next
    second = second.Next.Next
    if first == second{
      return true
    }
  }
  return false
}