package algorithms

// Quicksort order the array
func Quicksort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := arr[(left+right)/2]
	index := partition(arr, left, right, pivot)

	quicksort(arr, left, index-1)
	quicksort(arr, index, right)
}

func partition(arr []int, left, right, pivot int) int {
	for left <= right {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}
