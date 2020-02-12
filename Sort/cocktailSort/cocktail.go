package cocktailSort

func CocktailSort(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		left := 0
		right := len(arr) - 1
		for left <= right {
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++
			if arr[right-1] > arr[right] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--
		}
	}
	return arr
}
