package main

//基础版
func MergeSort(arr []int, n int) {
	doMergeSort(arr, 0, n-1)
}

func doMergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	doMergeSort(arr, l, mid)
	doMergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	aux := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}
	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

//优化版
func MergeSortOpt(arr []int, n int) {
	doMergeSortOpt(arr, 0, n-1)
}

func doMergeSortOpt(arr []int, l, r int) {
	//优化位置2.
	if r - l <= 15 {
		InsertionSortMerge(arr,l,r)
		return
	}
	mid := (l + r) / 2
	doMergeSortOpt(arr, l, mid)
	doMergeSortOpt(arr, mid+1, r)
	//优化位置1.
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r)
	}
}
//归并排序专用
func InsertionSortMerge(arr []int, l, r int) {
	for i := l+1; i <= r; i++ {
		t := arr[i]
		var j int
		for j = i; j > l && arr[j-1] > t; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
}
