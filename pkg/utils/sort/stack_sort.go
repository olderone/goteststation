package main

//最大堆
type MaxHeap struct {
	Data  []int
	count int
}

func (m *MaxHeap) MaxHeap(capacity int) {
	m.Data = make([]int, capacity)
	m.count = 0
}

func (m *MaxHeap) MaxHeap2(arr []int, n int){
	m.Data = make([]int, n+1)
	for i:=0;i<n;i++{
		m.Data = append(m.Data[:i+1],arr[i])
	}
	m.count = n
	count := m.count
	for k := count/2;k >= 1; k--{
		m.shiftDown(k)
	}
}

func (m *MaxHeap) Size() int {
	return m.count
}

func (m *MaxHeap) IsEmpty() bool {
	return m.count == 0
}

func (m *MaxHeap) Insert(data int) {
	m.Data = append(m.Data[:m.count+1],data)
	m.count++
	m.shiftUP()
}

func (m *MaxHeap) shiftUP() {
	d := m.count
	for d > 1 && m.Data[d] > m.Data[d/2] {
		m.Data[d], m.Data[d/2] = m.Data[d/2], m.Data[d]
		d = d/2
	}
}

func (m *MaxHeap)ExtractMax()int{
	if m.count <= 0{
		panic("index overflow")
	}

	t :=  m.Data[1]
	m.Data[1],m.Data[m.count] = m.Data[m.count],m.Data[1]
	m.count--
	m.shiftDown(1)

	return t
}

func (m *MaxHeap)shiftDown(k int){
	for k*2 <= m.count{
		j := k*2
		if j+1 <= m.count && m.Data[j+1] > m.Data[j]{
			j = j+1
		}
		if m.Data[k] > m.Data[j]{
			break
		}
		m.Data[k],m.Data[j] = m.Data[j],m.Data[k]
		k = j
	}
}

//堆排序1
func HeapSort(arr []int, n int){
	maxHeap := MaxHeap{}
	maxHeap.MaxHeap(n+1)
	for i := 0; i < n;i++{
		maxHeap.Insert(arr[i])
	}
	for i :=n-1;i >= 0 ; i--{
		arr = append(arr[:i],maxHeap.ExtractMax())
	}
}

//堆排序2
func HeapSort2(arr []int, n int){
	maxHeap := MaxHeap{}
	maxHeap.MaxHeap2(arr,n)
	for i :=n-1;i >= 0 ; i--{
		arr = append(arr[:i],maxHeap.ExtractMax())
	}
}

//堆排序3

func shiftDown(arr []int, n, i int){
	for i *2 +1 < n {
		j := i *2 +1
		if j+1 < n && arr[j+1] > arr[j]{
			j = j+1
		}

		if arr[i] > arr[j]{
			break
		}
		arr[i],arr[j] = arr[j],arr[i]
		i = j
	}
}

func HeapSort3(arr []int, n int){
	for i:= (n-1)/2; i >= 0 ; i--{
		shiftDown(arr,n,i)
	}
	for i := n-1; i > 0 ; i--{
		arr[i],arr[0] = arr[0],arr[i]
		shiftDown(arr,i,0)
	}
}