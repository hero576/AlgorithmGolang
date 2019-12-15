package array

import "fmt"

type ArrayListStack struct {
	array *ArrayList
}

func NewArrayListStack() *ArrayListStack {
	stack := new(ArrayListStack)
	stack.array = NewArrayList()
	return stack
}

func (s *ArrayListStack) Clear() {
	s.array.Clear()
}
func (s *ArrayListStack) Size() int {
	return s.array.Size()
}
func (s *ArrayListStack) Pop() interface{} {
	var (
		last interface{}
		err  error
	)
	if !s.IsEmpty() {
		if last, err = s.array.Get(s.array.Size() - 1); err != nil {
			fmt.Println(err)
		}
		if err = s.array.Delete(s.array.Size() - 1); err != nil {
			fmt.Println(err)
		}
		return last
	}
	return nil
}
func (s *ArrayListStack) Push(data interface{}) {
	if !s.IsFull() {
		s.array.Append(data)
	}
}
func (s *ArrayListStack) IsFull() bool {
	if s.array.Size() >= s.array.capacity {
		return true
	} else {
		return false
	}
}
func (s *ArrayListStack) IsEmpty() bool {
	if s.array.Size() == 0 {
		return true
	} else {
		return false
	}
}
