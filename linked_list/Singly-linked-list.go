package linked_list

import (
	"errors"
	"fmt"
)

type SinglyLinkedList struct {
	Head *Node
}

type Node struct {
	Data interface{}
	Next *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	head := &Node{0, nil}
	return &SinglyLinkedList{
		Head: head,
	}
}
func NewNode(data interface{}, next *Node) *Node {
	return &Node{Data: data, Next: next}
}

func (l *SinglyLinkedList) Len() int {
	length, ok := l.Head.Data.(int)
	if ok {
		return length
	}
	return 0
}

func (l *SinglyLinkedList) changeSize(incr int) {
	length := l.Len()
	l.Head.Data = length + incr
}

func (l *SinglyLinkedList) InsertAtFirst(node *Node) error {
	next := l.Head.Next
	node.Next = next
	l.Head.Next = node
	l.changeSize(1)
	return nil
}

func (l *SinglyLinkedList) Append(node *Node) error {
	current := l.Head
	for {
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	current.Next = node
	l.changeSize(1)
	return nil
}

func (l *SinglyLinkedList) Remove(data interface{}) error {
	if l.Len() == 0 {
		return errors.New("list is empty")
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			l.changeSize(-1)
			return nil
		}
		current = current.Next
	}
	return nil
}

func (l *SinglyLinkedList) Search(data interface{}) (*Node, error) {
	if l.Len() == 0 {
		return nil, errors.New("list is empty")
	}
	current := l.Head.Next
	if current.Data == data {
		return current, nil
	}
	for current.Next != nil {
		if current.Data == data {
			return current, nil
		}
		current = current.Next
	}
	if current.Data == data {
		return current, nil
	}
	return nil, errors.New("not find")
}

func (l *SinglyLinkedList) Range() {
	if l.Len() == 0 {
		fmt.Println("This is an empty list")
		return
	}
	fmt.Println("-------start----------","size:",l.Len())
	current := l.Head.Next
	i := 0
	for current.Next != nil {
		if i > 6{
			break
		}
		fmt.Printf("%+v \n",current)
		current = current.Next
		i++
	}
	fmt.Printf("%+v \n",current)
	fmt.Println("-------end----------")
	return

}

func (l *SinglyLinkedList)Reverse()  {
	if l.Len() == 0 {
		fmt.Println("This is an empty list")
		return
	}
	current := l.Head.Next

	for current.Next != nil {
		tmp := l.Head.Next
		new := current.Next
		current.Next = new.Next
		new.Next = tmp
		l.Head.Next = new
	}
}
