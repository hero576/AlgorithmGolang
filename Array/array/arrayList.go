package array

import (
	"AlgorithmGolang/Array/base"
	"errors"
	"fmt"
)

type ArrayList struct {
	data     []interface{}
	Capacity int
	size     int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.Capacity = capacity
	list.data = make([]interface{}, 0, list.Capacity)
	list.size = 0
	return list
}

func (list *ArrayList) Iterator() base.Iter {
	newit := new(ArrayListIterator)
	newit.currentIdex = 0
	newit.list = list
	return newit
}

func (list *ArrayList) Size() int {
	return list.size
}
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("index exceed")
	}
	return list.data[index], nil
}

func (list *ArrayList) checkFull() {
	if list.size == list.Capacity {
		newdataStore := make([]interface{}, 2*list.size, 2*list.size)
		copy(newdataStore, list.data)
		list.data = newdataStore
	}
}
func (list *ArrayList) Set(index int, value interface{}) error {
	if index < 0 || index >= list.size {
		return errors.New("index exceed")
	}
	list.data[index] = value
	return nil
}
func (list *ArrayList) Insert(index int, value interface{}) error {
	if index < 0 || index >= list.size {
		return errors.New("index exceed")
	}
	list.checkFull()
	list.data = append(list.data[:index], value, list.data[index:])

	list.size++
	return nil
}
func (list *ArrayList) Append(value interface{}) {
	list.data = append(list.data, value)
	list.size++
}
func (list *ArrayList) Clear() {
	list.data = make([]interface{}, 0, list.Capacity)
	list.size = 0
}
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.size {
		return errors.New("index exceed")
	}

	list.data = append(list.data[:index], list.data[index+1:]...)

	list.size--

	return nil
}
func (list *ArrayList) String() string {
	return fmt.Sprint(list.data)
}
