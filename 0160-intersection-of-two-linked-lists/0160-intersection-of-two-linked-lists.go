
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    firstA := headA
	secondB := headB
	for firstA != secondB{
		if firstA != nil {
			firstA = firstA.Next
		}else {
			firstA = headB
		}
		if secondB != nil {
			secondB = secondB.Next
		}else {
			secondB = headA
		}
	}
	return firstA
}