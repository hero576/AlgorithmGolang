package main

import (
	"AlgorithmGolang/Sort/bubbleSort"
	"AlgorithmGolang/Sort/cocktailSort"
	"AlgorithmGolang/Sort/heapSort"
	"AlgorithmGolang/Sort/insertSort"
	"AlgorithmGolang/Sort/mergeSort"
	"AlgorithmGolang/Sort/odd_evenSort"
	"AlgorithmGolang/Sort/redixSort"
	"AlgorithmGolang/Sort/selectSort"
	"AlgorithmGolang/Sort/shellSort"
	"AlgorithmGolang/treeSelectSort"
	"fmt"
)

func main() {
	arr := []int{1, 6, 2, 7, 23, 4, 7, 8}
	fmt.Println(selectSort.SelectSort(arr))
	arr2 := []string{"adf", "ad1", "adA", "fiogj", "b", "e", "123"}
	fmt.Println(selectSort.SelectSortString(arr2))

	fmt.Println(bubbleSort.BubbleSort(arr), "bubbleSort")
	fmt.Println(insertSort.InsertSort(arr), "insertSort")
	fmt.Println(heapSort.HeapSort(arr), "heapSort")
	fmt.Println(odd_evenSort.OddEvenSort(arr), "odd_evenSort")
	fmt.Println(mergeSort.MergeSort(arr), "mergeSort")
	fmt.Println(shellSort.ShellSort(arr), "shellSort")
	fmt.Println(redixSort.RedixSort(arr), "redixSort")
	fmt.Println(redixSort.RedixSortAdv(arr), "RedixSortAdv")
	fmt.Println(treeSelectSort.TreeSelectSort(arr), "TreeSelectSort")
	fmt.Println(cocktailSort.CocktailSort(arr), "CocktailSort")
}
