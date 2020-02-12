package main

import (
	"AlgorithmGolang/Array/array"
	"AlgorithmGolang/Array/queue"
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
	mystack := stack.NewArrayListStack()
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
	mystack := stack.NewArrayListIteratorStack()
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

func QueueArray1() {
	myqueue := queue.NewQueueArray()
	myqueue.EnQueue("1")
	myqueue.EnQueue("2")
	myqueue.EnQueue("3")
	myqueue.EnQueue("4")
	myqueue.EnQueue("5")
	fmt.Println(myqueue.DeQueue())
	fmt.Println(myqueue.DeQueue())
	fmt.Println(myqueue.DeQueue())
	fmt.Println(myqueue.DeQueue())
	fmt.Println(myqueue.DeQueue())
	fmt.Println(myqueue.DeQueue())
}
func CricleQuequArray() {
	q := queue.NewCircleQueue()
	q.EnQueue("1")
	q.EnQueue("1")
	q.EnQueue("1")
	q.EnQueue("a")
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())

}

func stackLink() {
	s := stack.NewStackLink()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	for data := s.Pop(); data != nil; data = s.Pop() {
		fmt.Println(data)
	}
}

func queueLink() {
	q := queue.NewQueueLink()
	q.EnQueue(1)
	q.EnQueue(1)
	q.EnQueue(1)
	q.EnQueue(2)
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}

func main() {
	//listIterator()
	//stackArray()
	//Arrayliststack()
	//QueueArray1()
	//CricleQuequArray()
	//stackLink()
	queueLink()
}
