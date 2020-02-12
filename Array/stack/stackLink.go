package stack

type StackLink struct {
	data  interface{}
	pNext *StackLink
}

func NewStackLink() *StackLink {
	return new(StackLink)
}

func (s *StackLink) Clear() {
	s = new(StackLink)
}

func (s *StackLink) Size() int {
	pNext := s
	length := 0
	for pNext.pNext != nil {
		pNext = pNext.pNext
		length++
	}
	return length
}
func (s *StackLink) Pop() interface{} {
	if s.IsEmpty() {
		return nil
		panic("stack is empty")
	}
	value := s.pNext.data
	s.pNext = s.pNext.pNext
	return value
}
func (s *StackLink) Push(data interface{}) {
	newStack := &StackLink{data: data}
	//头部插入
	newStack.pNext = s.pNext
	s.pNext = newStack
}
func (s *StackLink) IsFull() bool {
	return false
}
func (s *StackLink) IsEmpty() bool {
	return s.pNext == nil
}
