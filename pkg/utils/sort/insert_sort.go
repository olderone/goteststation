package main

//基础版
func InsertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
}

//优化版
func InsertionSortNew(arr []int, n int) {
	for i := 1; i < n; i ++ {
		var j int
		t := arr[i]
		for j = i; j > 0 && arr[j-1] > t; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
}
