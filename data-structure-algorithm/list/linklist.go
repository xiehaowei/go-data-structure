package list

import "fmt"

type LinkList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func GetNode(data int) *Node {
	return &Node{Data: data}
}

func InitLinkList(param ...int) LinkList {
	head := new(Node)
	cur := head
	for _, one := range param {
		tmp := &Node{
			Data: one,
		}
		cur.Next = tmp
		cur = tmp
	}
	return LinkList{Head: head}
}

func (l LinkList) PrintLinkList() {
	curr := l.Head
	if curr.Next == nil {
		fmt.Println("empty linklist")
		return
	}
	for curr.Next != nil {
		fmt.Println(curr.Next.Data)
		curr = curr.Next
	}
}

func (l *LinkList) InsertOne(data int) {
	curr := l.Head
	if curr.Next == nil {
		fmt.Println("empty linklist")
		return
	}
	for curr.Next != nil {
		if curr.Next.Data > data {
			node := &Node{
				Data: data,
				Next: curr.Next,
			}
			curr.Next = node
			return
		} else {
			curr = curr.Next
		}

	}
}

func (l *LinkList) DeleteOne(data int) {
	curr := l.Head
	if curr.Next == nil {
		fmt.Println("empty linklist")
		return
	}
	for curr.Next != nil {
		if curr.Next.Data == data {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}
}

func (l *LinkList) AppendOne(data int) {
	curr := l.Head
	if curr.Next == nil {
		fmt.Println("empty linklist")
		return
	}
	if curr.Next.Next == nil {
		curr.Next.Next = GetNode(data)
		return
	}
	for curr.Next != nil && curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next.Next = GetNode(data)
}

//
func (l *LinkList) Reverse() {
	h := l.Head
	if h.Next == nil {
		fmt.Println("empty linklist")
		return
	}
	current := h.Next
	h.Next = nil //重要！！！
	for current != nil {
		fmt.Println(current.Data)
		n := current.Next
		tmpF := h.Next //头节点后第一个
		current.Next = tmpF
		h.Next = current
		current = n
	}
}
func (l *LinkList) HasCyrle() bool {
	if l.Head.Next == nil {
		return false
	}
	slow := l.Head.Next
	fast := l.Head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func Merge2SortedList(l1, l2 LinkList) LinkList {
	if nil == l1.Head || nil == l1.Head.Next {
		return l2
	}
	if nil == l2.Head || nil == l2.Head.Next {
		return l1
	}
	l := LinkList{Head: &Node{}}
	cur := l.Head
	cur1 := l1.Head.Next
	cur2 := l2.Head.Next

	for cur1 != nil && cur2 != nil {
		if cur1.Data > cur2.Data {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}
	return l
}

func (l *LinkList) DeleteBottomN(n int) {
	if n < 0 || l.Head == nil || l.Head.Next == nil {
		return
	}

	fast := l.Head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return
	}
	slow := l.Head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

}

func (l *LinkList) FindMiddleNode() *Node {
	if nil == l.Head || nil == l.Head.Next {
		return nil
	}
	if nil == l.Head.Next.Next {
		return l.Head.Next
	}

	slow, fast := l.Head, l.Head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
