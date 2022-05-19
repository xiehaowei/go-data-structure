package list

import (
	"fmt"
	"testing"
)

func Test_InitLinkList(t *testing.T) {
	l := InitLinkList(1, 2, 3, 5, 8, 13)
	l.PrintLinkList()
	l.Reverse()
	fmt.Println("--------------")
	l.PrintLinkList()
}

func Test_PrintLinkList(t *testing.T) {
	l := LinkList{Head: &Node{}}
	l.PrintLinkList()
	fmt.Println("--------------")
	l = InitLinkList(1, 2, 3, 5, 8, 13)
	l.PrintLinkList()
	fmt.Println("--------------")
	l.InsertOne(10)
	l.PrintLinkList()
	fmt.Println("--------------")
	l.DeleteOne(10)
	l.PrintLinkList()
}

func Test_Lru(t *testing.T) {
	lru := NewLRU(5)
	lru.Visit(1)
	lru.Visit(2)
	lru.Visit(3)
	lru.Visit(4)
	lru.Print()
	lru.Visit(5)
	lru.Print()
	lru.Visit(4)
	fmt.Println("visit 4")
	lru.Print()
	lru.Visit(6)
	fmt.Println("visit 6")
	lru.Print()
	lru.Visit(3)
	fmt.Println("visit 3")
	lru.Print()
	lru.Visit(4)
	fmt.Println("visit 4")
	lru.Print()
}

func Test_HasCyrle(t *testing.T) {
	l := InitLinkList(1, 2, 3, 5, 8, 13)
	fmt.Println(l.HasCyrle())

	n1 := GetNode(1)
	l2 := LinkList{
		Head: &Node{
			Next: n1,
		},
	}
	n2 := GetNode(2)
	n3 := GetNode(3)
	n4 := GetNode(4)
	n5 := GetNode(5)
	n6 := GetNode(6)
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n2
	fmt.Println(l2.HasCyrle())

}

func Test_MergeSortedList(t *testing.T) {
	l1 := InitLinkList(1, 2, 3, 5, 8, 13)
	l2 := InitLinkList(4, 6, 10, 12, 14, 16, 18, 20, 22)
	l3 := Merge2SortedList(l1, l2)
	l3.PrintLinkList()
}

func Test_DeleteBottomN(t *testing.T) {
	l1 := InitLinkList(1, 2, 3, 5, 8, 13)
	l1.DeleteBottomN(1)
	l1.PrintLinkList()
}

func Test_FindMiddle(t *testing.T) {
	l1 := InitLinkList(1, 2, 3, 5, 8, 13)
	node := l1.FindMiddleNode()
	fmt.Println(*node)
	l1.AppendOne(21)
	l1.PrintLinkList()
	node = l1.FindMiddleNode()
	fmt.Println(*node)
}
