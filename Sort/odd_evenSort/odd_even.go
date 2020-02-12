package odd_evenSort

func OddEvenSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		isSord := false
		for isSord == false {
			isSord = true
			for i := 1; i < length-1; i++ {
				if arr[i] < arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
					isSord = false
				}
			}
			for i := 0; i < length-1; i++ {
				if arr[i] < arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
					isSord = false
				}
			}
		}
		return arr
	}

}
