package redixSort

import (
	"AlgorithmGolang/Sort/selectSort"
)

//基数排序或者叫做桶排序

func RedixSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		max := selectSort.SelectMax(arr)
		for bit := 1; max/bit > 10; bit *= 10 {

			bitCount := make([]int, 10)
			for k := 0; k < length; k++ {
				num := (arr[k] / bit) % 10
				bitCount[num]++
			}
			for k := 1; k < 10; k++ {
				bitCount[k] += bitCount[k-1]
			}
			tmp := make([]int, 10)
			for k := length - 1; k >= 0; k-- {
				num := (arr[k] / bit) % 10
				tmp[bitCount[num]-1] = arr[k]
				bitCount[num]--
			}
			for k := 0; k < length; k++ {
				arr[k] = tmp[k]
			}
		}
		return arr
	}
}

func RedixSortAdv(arr []int) []int {
	//计数排序
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		max := selectSort.SelectMax(arr)
		sortedArr := make([]int, len(arr))
		countsArr := make([]int, max+1)
		for _, v := range arr {
			countsArr[v]++
		}
		for i := 1; i <= max; i++ {
			countsArr[i] += countsArr[i-1]
		}
		for _, v := range arr {
			sortedArr[countsArr[v]-1] = v
			countsArr[v]--
		}

		return sortedArr
	}
}
