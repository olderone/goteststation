package main

//基础版
func BubbleSort(arr []int, n int) {
	for {
		newN := true
		for i := 1; i < n; i ++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				newN = false
			}
		}
		n--
		if newN {
			break
		}
	}
}

//优化版
func BubbleSortNew(arr []int, n int) {
	for {
		newN := 0
		for i := 1; i < n; i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				newN = i
			}
		}
		n = newN
		if newN == 0 {
			break
		}
	}
}
