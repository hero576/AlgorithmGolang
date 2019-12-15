package main

import (
	"AlgorithmGolang/Array/array"
	"fmt"
)

func main() {
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
