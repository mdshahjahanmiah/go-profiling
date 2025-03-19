package stack

type LinkedList struct {
	top *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (s *LinkedList) Push(value int) {
	newNode := &Node{value: value, next: s.top} // Heap allocation for each node
	s.top = newNode
}

func (s *LinkedList) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	value := s.top.value
	s.top = s.top.next // No heap reallocation, just pointer reassignment
	return value, true
}
