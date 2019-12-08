package array

import (
	"AlgorithmGolang/Array/base"
	"errors"
)

type ArrayListIter interface {
	Iterator() base.Iter
}

type ArrayListIterator struct {
	list        *ArrayList
	currentIdex int
}

func (list *ArrayList) Iterable() base.Iter {
	it := new(ArrayListIterator)
	it.currentIdex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("has not next")
	}
	value, err := it.list.Get(it.currentIdex)
	it.currentIdex++
	return value, err
}
func (it *ArrayListIterator) Remove() (err error) {
	it.currentIdex--
	err = it.list.Delete(it.currentIdex) //由于迭代器首先迭代到当前元素，currentindex指向了下一个，所以删除当前元素就需要先--
	return
}
func (it *ArrayListIterator) GetIndex() int {
	return it.currentIdex
}
func (it *ArrayListIterator) HasNext() bool {
	return it.currentIdex < it.list.Size()
}
