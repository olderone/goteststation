package main

func MergeSortBU(arr []int, n int) {
	for sz := 1; sz < n; sz += sz {
		for i := sz; i < n-sz; i += sz + sz {
			merge(arr, i, i+sz-1, min(i+sz+sz-1, n-1))
		}
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
