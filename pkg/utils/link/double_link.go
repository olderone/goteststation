package main

type DoubleNode struct {
	data int
	Next *DoubleNode
	Prev *DoubleNode
}

type DoubleList struct {
	Head *DoubleNode
	Tail *DoubleNode
	Size int
}

func (l *DoubleList) TailAdd(data int) {
	newP := new(DoubleNode)
	newP.data = data
	if l.Tail != nil {
		newP.Prev = l.Tail
		l.Tail.Next = newP
	} else {
		l.Head = newP
	}
	l.Tail = newP
	l.Size ++
	return
}

func (l *DoubleList) HeadAdd(data int) {
	newP := new(DoubleNode)
	newP.data = data
	if l.Head != nil {
		newP.Next = l.Head
		l.Head.Prev = newP
	} else {
		l.Tail = newP
	}
	l.Head = newP
	l.Size++
	return

}

func (l *DoubleList) Remove(node *DoubleNode) bool {
	if node == nil {
		return false
	}
	prev := node.Prev
	next := node.Next
	if node == l.Head {
		l.Head = next
	} else {
		prev.Next = next
	}
	if node == l.Tail {
		l.Tail = prev
	} else {
		next.Prev = prev
	}
	l.Size --
	return true
}

func (l *DoubleList) SizeList() int {
	return l.Size
}