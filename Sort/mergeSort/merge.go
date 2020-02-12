package mergeSort

import "AlgorithmGolang/Sort/insertSort"

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else if length > 1 && length < 5 { //加速
		return insertSort.InsertSort(arr)
	} else {
		mid := length / 2

		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])

		//归并
		indexLeft := 0
		indexRight := 0
		arrRes := []int{}
		for indexLeft < len(leftarr) && indexRight < len(rightarr) {
			if leftarr[indexLeft] > rightarr[indexRight] {
				arrRes = append(arrRes, leftarr[indexLeft])
				indexLeft++
			} else {
				arrRes = append(arrRes, rightarr[indexRight])
				indexRight++
			}
		}
		for indexLeft < len(leftarr) {
			arrRes = append(arrRes, leftarr[indexLeft])
			indexLeft++
		}
		for indexRight < len(rightarr) {
			arrRes = append(arrRes, rightarr[indexRight])
			indexRight++
		}
		return arrRes
	}
}
