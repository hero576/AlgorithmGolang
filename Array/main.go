package main

import (
	"AlgorithmGolang/Array/array"
	"AlgorithmGolang/Array/stack"
	"fmt"
)

func listIterator() {
	var list *array.ArrayList = array.NewArrayList()
	list.Append("1")
	list.Append("2")
	list.Append("3")
	list.Append("a")
	list.Append("b")
	for it := list.Iterator(); it.HasNext(); {
		if item, err := it.Next(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(item)
		}
	}
}

func stackArray() {
	mystack := stack.NewStackArray()
	mystack.Push("1")
	mystack.Push("2")
	mystack.Push("3")
	mystack.Push("a")
	mystack.Push("b")
	for value := mystack.Pop(); value != nil; value = mystack.Pop() {
		fmt.Println(value)
	}
}

func Arrayliststack() {
	mystack := array.NewArrayListStack()
	mystack.Push("1")
	mystack.Push("2")
	mystack.Push("3")
	mystack.Push("a")
	mystack.Push("b")
	for value := mystack.Pop(); value != nil; value = mystack.Pop() {
		fmt.Println(value)
	}
}

func arraylistIteratorStack() {
	mystack := array.NewArrayListIteratorStack()
	mystack.Push("1")
	mystack.Push("2")
	mystack.Push("3")
	mystack.Push("a")
	mystack.Push("b")
	for iter := mystack.It; iter.HasNext(); {
		if item, err := iter.Next(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(item)
		}
	}
}

func main() {
	//listIterator()
	//stackArray()
	Arrayliststack()
}
