package stack

type StackArray struct {
	dataSource  []interface{}
	capSize     int
	currentSize int
}

func NewStackArray() *StackArray {
	stack := new(StackArray)
	stack.dataSource = make([]interface{}, 0, 10)
	stack.capSize = 10
	stack.currentSize = 0
	return stack
}

func (s *StackArray) Clear() {
	s.dataSource = make([]interface{}, 0, 10)
	s.capSize = 10
	s.currentSize = 0
}
func (s *StackArray) Size() int {
	return s.currentSize
}
func (s *StackArray) Pop() interface{} {
	if !s.IsEmpty() {
		s.currentSize--
		last := s.dataSource[s.currentSize]
		s.dataSource = s.dataSource[:s.currentSize]
		return last
	}
	return nil
}
func (s *StackArray) Push(data interface{}) {
	if !s.IsFull() {
		s.dataSource = append(s.dataSource, data)
		s.currentSize++
	}
}
func (s *StackArray) IsFull() bool {
	if s.currentSize >= s.capSize {
		return true
	} else {
		return false
	}
}
func (s *StackArray) IsEmpty() bool {
	if s.currentSize == 0 {
		return true
	} else {
		return false
	}
}
