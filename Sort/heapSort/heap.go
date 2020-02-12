package heapSort

func HeapSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1
		for i := 0; i < length-1; i++ {
			arrCopy := arr[i:]
			lengthCopy := len(arrCopy)

			for j := depth; j >= 0; j-- {
				topmax := j
				leftChild := 2*j + 1
				rightChild := 2*j + 2
				if leftChild < lengthCopy-1 && arrCopy[leftChild] > arrCopy[topmax] {
					topmax = leftChild
				}
				if rightChild < lengthCopy-2 && arrCopy[rightChild] > arrCopy[topmax] {
					topmax = rightChild
				}
				if topmax != j {
					arrCopy[j], arrCopy[topmax] = arrCopy[topmax], arrCopy[j]
				}
			}
			arr = append(arr[:i], arrCopy...)
		}

		return arr
	}

}
