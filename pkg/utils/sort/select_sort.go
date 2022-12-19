package main

//基础版
func SelectionSort(arr []int, n int) {
	for i := 0; i < n; i ++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
}

//优化版本
func SelectionSortNew(arr []int, n int) {
	left := 0
	right := n - 1
	for left < right {
		minIndex := left
		maxIndex := right

		if arr[minIndex] > arr[maxIndex] {
			arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]
		}

		for i := left + 1; i < right; i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			} else if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}

		arr[minIndex], arr[left] = arr[left], arr[minIndex]
		arr[maxIndex], arr[right] = arr[right], arr[maxIndex]
		left++
		right++
	}
}