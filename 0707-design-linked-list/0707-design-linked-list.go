type node struct {
	data int
	next *node
}
type MyLinkedList struct {
	head *node
	size int
}


func Constructor() MyLinkedList {
	list := MyLinkedList{}
	return list
}

func (thisList *MyLinkedList) Get(index int) int {
	if index < 0 || index >= thisList.size {
		return -1
	}
	if thisList.head == nil {
		return -1
	}

	pointer := thisList.head
	for i := 0; i < index; i++ {
		pointer = pointer.next
	}
	return pointer.data
}

func (thisList *MyLinkedList) AddAtHead(val int) {
	headNode := &node{val, nil}
	headNode.next = thisList.head
	thisList.head = headNode
	thisList.size++
}
func (thisList *MyLinkedList) AddAtTail(val int) {
	tailNode := thisList.head
	if tailNode == nil {
		thisList.head = &node{val, nil}
	} else {
		for tailNode.next != nil {
			tailNode = tailNode.next
		}
		tailNode.next = &node{val, nil}
	}
	thisList.size++
}

func (thisList *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > thisList.size {
		return
	}
	if index == 0 {
		thisList.AddAtHead(val)
	} else {
		currentNode := thisList.head

		for i := 0; i < index-1; i++ {
			currentNode = currentNode.next
		}
		newNode := &node{val, nil}
		newNode.next = currentNode.next
		currentNode.next = newNode
		thisList.size++
	}
}

func (thisList *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= thisList.size {
		return
	}
	pointer := thisList.head
	if index == 0 {
		thisList.head = pointer.next
	} else {
		for i := 0; i < index-1; i++ {
			pointer = pointer.next
		}
		pointer.next = pointer.next.next
		thisList.size--
	}
}
