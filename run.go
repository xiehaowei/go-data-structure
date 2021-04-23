package main

import (
	"fmt"
	"github.com/xiehaowei/go-data-structure/linked_list"
	"github.com/xiehaowei/go-data-structure/stack"
)

func main() {
	testLinkedList()
	//testStack()
}

func testStack() {
	s := stack.NewStack()
	fmt.Println("len : ", s.Len())
	fmt.Println(s.Pop())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println("len : ", s.Len())
	s.PrintAll()
	fmt.Println(s.Pop())
	fmt.Println("len : ", s.Len())

	s.PrintAll()

}

func testLinkedList() {
	l := linked_list.NewSinglyLinkedList()
	node1 := linked_list.NewNode(1, nil)
	node2 := linked_list.NewNode(2, nil)
	node0 := linked_list.NewNode(0, nil)
	node3 := linked_list.NewNode(3, nil)
	node4 := linked_list.NewNode(4, nil)
	node5 := linked_list.NewNode(5, nil)
	l.Append(node1)
	l.Append(node2)
	l.Append(node3)
	l.Append(node4)
	l.Append(node5)
	l.InsertAtFirst(node0)
	l.Range()
	l.Reverse()
	l.Range()
}
