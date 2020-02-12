package shellSort

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				//使用插入排序
				for k := 1; k < length; k += gap {
					backup := arr[k]
					j := k - 1
					for j >= 0 && backup > arr[j] {
						arr[j+1] = arr[j]
						j--
					}
					arr[j+1] = backup
				}
			}
			gap /= 2 //或者gap--
		}
		return arr
	}
}
