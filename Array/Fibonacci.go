package main

import (
	"AlgorithmGolang/Array/stack"
	"errors"
	"fmt"
)

func Fib(num int) int {
	if num <= 0 {
		panic(errors.New("num cannot be below zero"))
	}
	if num == 1 || num == 2 {
		return 1
	}
	return Fib(num-1) + Fib(num-2)
}

func FibStack(num int) int {
	var (
		mystack *stack.StackArray
		data    interface{}
		last    int
	)
	if num <= 0 {
		panic(errors.New("num cannot be below zero"))
	}
	if num == 1 || num == 2 {
		return 1
	}
	mystack = stack.NewStackArray()
	mystack.Push(1)
	last = 0                    //保存最终结果
	for i := 0; i <= num; i++ { // 判断堆栈不为空
		data = mystack.Pop()
		if i == 0 || i == 1 {
			last = 1
			mystack.Push(1)
		} else {
			mystack.Push(last + data.(int))
			last = data.(int)
		}
	}
	return last
}

func main() {
	//fmt.Println(Fib(6))
	for i := 1; i < 10; i++ {
		fmt.Println(FibStack(i))
	}
}
