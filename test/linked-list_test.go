package test

import (
	"fmt"
	"github.com/xiehaowei/go-data-structure/linked_list"
	"sync"
	"testing"
)

var l *linked_list.SinglyLinkedList
var once sync.Once

//func TestLinkedListAppend(t *testing.T) {
//	l2 := linked_list.NewSinglyLinkedList()
//	l2.Range()
//	node1 := linked_list.NewNode(1, nil)
//	node2 := linked_list.NewNode(2, nil)
//	err := l2.Append(node1)
//	if err != nil {
//		t.Fail()
//	}
//	err = l2.Append(node2)
//	if err != nil {
//		t.Fail()
//	}
//	if l2.Len() != 2 {
//		t.Fail()
//	}
//}

func Benchmark_LinkedListAppend(b *testing.B) {
	b.StopTimer()
	l = linked_list.NewSinglyLinkedList()
	for i:=1;i<= 99999;i++{
		node := linked_list.NewNode(i, nil)
		l.Append(node)
	}
	b.StartTimer()
	fmt.Println("append l.size = ",l.Len())
	node := linked_list.NewNode(999999999999, nil)
	l.Append(node)
}

func Benchmark_LinkedListSearch(b *testing.B) {
	b.StopTimer()
	l = linked_list.NewSinglyLinkedList()
	for i:=1;i<= 99999;i++{
		node := linked_list.NewNode(i, nil)
		l.Append(node)
	}
	b.StartTimer()
	fmt.Println("search l.size = ",l.Len())
	_, err := l.Search(99998)
	if err != nil {
		b.Fail()
	}
}

func getList(b *testing.B) *linked_list.SinglyLinkedList {
	once.Do(func() {
		l = linked_list.NewSinglyLinkedList()
	})
	for i:=1;i<= b.N;i++{
		node := linked_list.NewNode(i, nil)
		l.Append(node)
	}
	return l
}
