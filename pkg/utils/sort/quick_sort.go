package main

import (
	"fmt"
	"math/rand"
)

//基础版
func quickSort(arr []int, n int) {
	doQuickSort(arr, 0, n-1)
}

func doQuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := doPartition(arr, l, r)
	doQuickSort(arr, l, p-1)
	doQuickSort(arr, p+1, r)
}

func doPartition(arr []int, l, r int) int {
	v := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			arr[i], arr[j+1] = arr[j+1], arr[i]
			j++
		}
	}
	arr[l],arr[j] = arr[j],arr[l]
	return j
}

//优化版
func quickSortOpt(arr []int, n int) {
	doQuickSortOpt(arr, 0, n-1)
}

func doQuickSortOpt(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := doPartitionOpt(arr, l, r)
	doQuickSortOpt(arr, l, p-1)
	doQuickSortOpt(arr, p+1, r)
}

func doPartitionOpt(arr []int, l, r int) int {
	//优化位置
	t := rand.Intn(r)%(r-l+1) + l
	arr[l], arr[t] = arr[t], arr[l]
	v := arr[l]

	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			arr[i], arr[j+1] = arr[j+1], arr[i]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//双路优化版
func QuickSort2(arr []int, n int) {
	doQuickSort2(arr, 0, n-1)
}

func doQuickSort2(arr []int, r int) {
	if r-l <= 15 {
		InsertionSortMerge(arr, l, r)
		return
	}

	p := partition2(arr, l, r)
	doQuickSort2(arr, l, p-1)
	doQuickSort2(arr, p+1, r)
}

func partition2(arr []int, l, r int) int {
	t := rand.Intn(r)%(r-l+1) + l
	arr[l], arr[t] = arr[t], arr[l]
	v := arr[l]

	i := l + 1
	j := r
	for {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[i],arr[j] = arr[j],arr[i]
		i++
		j--
	}
	arr[l],arr[j] = arr[j],arr[l]
	return j
}

//三路优化版
func quickSort3ways(arr []int, n int) {
	doQuickSort3ways(arr, 0, n-1)
}

func doQuickSort3ways(arr []int, l, r int) {
	if r-l <= 15 {
		InsertionSortMerge(arr, l, r)
		return
	}

	lt, gt := partition3ways(arr, l, r)
	doQuickSort3ways(arr, l, lt-1)
	doQuickSort3ways(arr, gt, r)
}

func partition3ways(arr []int, l, r int) (ltd, gtd int) {
	t := rand.Intn(r)%(r-l+1) + l
	arr[l], arr[t] = arr[t], arr[l]
	v := arr[l]

	lt := l
	gt := r + 1
	i := l+1
	for i < gt{
		if arr[i] < v{
			arr[i],arr[lt+1] = arr[lt+1],arr[i]
			i++
			lt++
		}else if arr[i] >v {
			arr[i],arr[gt-1] = arr[gt-1],arr[i]
			gt--
		}else{
			i++
		}
	}
	arr[l],arr[lt] = arr[lt],arr[l]
	return lt,gt
}


func FindArray(arr []int, n int) {
	doFindArray(arr, 0, n-1)
}

func doFindArray(arr []int, l, r int)int{
	if l >= r {
		return arr[l]
	}

	lt, gt := partitionFindArray(arr, l, r)
	if lt ==1000 {
		fmt.Println(arr[1000])
		return arr[1000]
	}else if 1000 < lt{
		return doFindArray(arr, l, lt-1)
	}else{
		return doFindArray(arr, gt, r)
	}
}

func partitionFindArray(arr []int, l, r int) (int, int) {
	t := rand.Intn(r)%(r-l+1) + l
	arr[l], arr[t] = arr[t], arr[l]
	v := arr[l]

	lt := l
	gt := r + 1
	i := l + 1
	for i < gt {
		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			i++
			lt++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	return lt, gt
}