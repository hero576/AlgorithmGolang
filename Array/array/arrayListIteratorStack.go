package array

import (
	"AlgorithmGolang/Array/base"
	"fmt"
)

type ArrayListIteratorStack struct {
	array *ArrayList
	It    base.Iter
}

func NewArrayListIteratorStack() *ArrayListIteratorStack {
	it := new(ArrayListIteratorStack)
	it.array = NewArrayList()
	it.It = it.array.Iterator()
	return it
}

func (s *ArrayListIteratorStack) Clear() {
	s.array.Clear()
}
func (s *ArrayListIteratorStack) Size() int {
	return s.array.Size()
}
func (s *ArrayListIteratorStack) Pop() interface{} {
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
func (s *ArrayListIteratorStack) Push(data interface{}) {
	if !s.IsFull() {
		s.array.Append(data)
	}
}
func (s *ArrayListIteratorStack) IsFull() bool {
	if s.array.Size() >= s.array.capacity {
		return true
	} else {
		return false
	}
}
func (s *ArrayListIteratorStack) IsEmpty() bool {
	if s.array.Size() == 0 {
		return true
	} else {
		return false
	}
}
