package bubbleSort

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			isNeedExchange := false // 不需要交换提前退出
			for j := 0; j < length-1-i; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					isNeedExchange = true
				}
			}
			if !isNeedExchange {
				break
			}
		}
		return arr
	}
}
