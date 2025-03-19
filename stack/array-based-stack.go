package stack

type Array struct {
	items []int
}

func NewArray(capacity int) *Array {
	return &Array{
		items: make([]int, 0, capacity), // Pre-allocate to avoid resizing
	}
}

func (s *Array) Push(value int) {
	s.items = append(s.items, value) // Appends to the slice, potentially triggering heap allocation
}

func (s *Array) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index] // Shrinks the slice, potentially triggering GC
	return value, true
}
