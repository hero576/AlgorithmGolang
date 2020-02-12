package treeSelectSort

import (
	"fmt"
	"math"
)

type node struct {
	value   int
	isValid bool
	rank    int
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func TreeSelectSort(arr []int) []int {
	level := 0
	result := make([]int, 0, len(arr))
	for pow(2, level) < len(arr) {
		level++
	}
	leaf := pow(2, level)
	tree := make([]node, leaf*2)

	for i := 0; i < len(arr); i++ {
		tree[leaf+i] = node{arr[i], true, i}
	}
	fmt.Println(tree)
	for i := 0; i < level; i++ {
		nodeCount := pow(2, level)
		for j := 0; j < nodeCount/2; j++ {
			leftnode := nodeCount + j*2
			rightnode := leftnode + 1
			fmt.Println(leftnode, rightnode, nodeCount)
			if !tree[leftnode].isValid || tree[rightnode].isValid && (tree[leftnode].value > tree[rightnode].value) {
				mid := (leftnode - 1) / 2
				tree[mid] = tree[rightnode]
				fmt.Println("tree[mid].value1", mid, tree[mid].value, tree[rightnode], rightnode)
			} else {
				mid := (leftnode - 1) / 2
				tree[mid] = tree[leftnode]
				fmt.Println("tree[mid].value2", mid, tree[mid].value, tree[leftnode], leftnode)
			}
		}
		//todo cannot insert value
		fmt.Println("tree0", tree[0])
		result = append(result, tree[0].value)
	}

	for t := 0; t < len(arr)-1; t++ {
		winNode := tree[0].rank + leaf - 1
		tree[winNode].isValid = false
		for i := 0; i < level; i++ {
			leftNode := winNode
			if winNode%2 == 0 {
				leftNode = winNode - 1
			}
			rightNode := leftNode + 1

			if !tree[leftNode].isValid || tree[rightNode].isValid && (tree[leftNode].value > tree[rightNode].value) {
				mid := (leftNode - 1) / 2
				tree[mid].value = tree[rightNode].value
			} else {
				mid := (leftNode - 1) / 2
				tree[mid].value = tree[leftNode].value
			}

			winNode = (leftNode - 1) / 2
		}
		result = append(result, tree[0].value)
	}
	fmt.Println(result)
	return arr
}
