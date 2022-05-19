package list

import (
	"fmt"
)

type NodeDoubleDirect struct {
	Data int
	Next *NodeDoubleDirect
}
type LRU struct {
	Cap  int
	Next *NodeDoubleDirect
	Size int
}

func NewLRU(size int) LRU {
	return LRU{
		Cap: size,
	}
}

func (l *LRU) Print() {
	fmt.Println("current cap is ", l.Size)
	node := l.Next
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}

func (l *LRU) Visit(data int) {
	node := l.Next
	if node == nil {
		new := &NodeDoubleDirect{
			Data: data,
		}
		l.Next = new
		l.Size++
		return
	}

	if l.Size == 1 && node.Data == data {
		return
	}

	prev := node
	prevv := node
	for node != nil {
		if node.Data == data {
			nn := node.Next
			prev.Next = nn
			f := l.Next
			node.Next = f
			l.Next = node
			return
		}
		prev = node
		if node.Next != nil && node.Next.Next == nil {
			prevv = node
		}
		node = node.Next
	}

	//直接插入头部

	first := l.Next
	new := &NodeDoubleDirect{
		Data: data,
		Next: first,
	}
	l.Next = new

	if l.Size < l.Cap {
		l.Size++
	} else {
		prevv.Next = nil
	}
}
