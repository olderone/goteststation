package main

import "fmt"

type Node struct {
	data int
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) TailList(data int, p *Node){
	newP := new(Node)
	newP.data = data
	newP.Next = nil
	if l.Head == nil {
		l.Head = newP
		return
	}
	//p := l.Head
	if p == nil {
		p = l.Head
	}
	for p.Next != nil {
		p = p.Next
	}
	p.Next = newP
	fmt.Println("p:", p)
	return
}

func (l *List) HeadList(data int) {
	newP := new(Node)
	newP.data = data
	newP.Next = nil
	if l.Head == nil {
		l.Head = newP
		return
	}
	newP.Next = l.Head
	l.Head = newP
	return
}

func (l *List) Delete(data int) {
	if l.Head == nil {
		return
	}
	p := l.Head
	//如果是表头
	if p.data == data {
		l.Head = p.Next
		return
	}
	p1 := p
	for p.Next != nil {
		if p.data == data {
			p1.Next = p.Next
			return
		}
		fmt.Printf("p.data:%d ", p.data)
		p1 = p
		fmt.Printf("p1.data:%d ", p1.data)
		p = p.Next
	}
	if p.data == data {
		p1.Next = nil
		return
	}
}
func main() {
	l := List{}
	t := []int{1, 2, 3, 4, 5, 6, 7, 8}
	lb := l.Head
	for _, v := range t {
		l.TailList(v, lb)
	}
	fmt.Println("lb:", fmt.Sprintf("%T", l.Head.data), lb)
	p := l.Head
	for p != nil {
		fmt.Printf("p.data:%d ", p.data)
		p = p.Next
	}
	l.Delete(8)
	fmt.Println()
	p = l.Head
	for p != nil {
		fmt.Printf("p.data:%d ", p.data)
		p = p.Next
	}

}