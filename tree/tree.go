package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func CreateNode(v int) *Node {
	return &Node{
		Value: v,
	}
}

func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, ",")
	PreOrder(node.Left)
	PreOrder(node.Right)
}
func MidOrder(node *Node) {
	if node == nil {
		return
	}
	MidOrder(node.Left)
	fmt.Print(node.Value, ",")
	MidOrder(node.Right)
}
func PostOrder(node *Node) {
	if node == nil {
		return
	}

	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Print(node.Value, ",")
}

func PreOrderNotRecursion(root *Node) {
	vals := make([]int, 0, 5)
	stack := []*Node{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Value)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	for _, v := range vals {
		fmt.Print(v, ",")
	}
}

func InOrderNotRecursion(root *Node) {
	vals := make([]int, 0, 5)
	stack := []*Node{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		vals = append(vals, node.Value)
		node = node.Right
		stack = stack[:len(stack)-1]
	}
	for _, v := range vals {
		fmt.Print(v, ",")
	}
}

func PostOrderNotRecursion(root *Node) {
	vals := make([]int, 0, 5)
	stack := []*Node{}
	node := root
	var last *Node
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if node.Right == nil || node.Right == last {
			vals = append(vals, node.Value)
			last = node
			node = nil
			stack = stack[:len(stack)-1]
		} else {
			node = node.Right
		}
	}
	for _, v := range vals {
		fmt.Print(v, ",")
	}
}

func BreadthFirstSearch(root *Node) {
	result := make([]int, 0, 5)
	nodeList := make([]*Node, 0, 5)
	nodeList = append(nodeList, root)
	for len(nodeList) > 0 {
		node := nodeList[0]
		result = append(result, node.Value)
		nodeList = nodeList[1:]
		if node.Left != nil {
			nodeList = append(nodeList, node.Left)
		}
		if node.Right != nil {
			nodeList = append(nodeList, node.Right)
		}
	}
	for _, v := range result {
		fmt.Print(v, ",")
	}
}

/**
				1
            2      3
         4    5       6
                 7
               8
*/
func main() {
	tree := CreateNode(1)
	n2 := CreateNode(2)
	n3 := CreateNode(3)
	n4 := CreateNode(4)
	n5 := CreateNode(5)
	n6 := CreateNode(6)
	n7 := CreateNode(7)
	n8 := CreateNode(8)
	tree.Left = n2
	tree.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Right = n6
	n5.Left = n7
	n7.Right = n8
	fmt.Println("preOrder")
	PreOrder(tree)
	fmt.Println("\n ============ ")

	fmt.Println("PreOrderNotRecursion")
	PreOrderNotRecursion(tree)
	fmt.Println("\n ============ ")

	fmt.Println("midOrder")
	MidOrder(tree)
	fmt.Println("\n ============ ")

	fmt.Println("inOrderNotRecursion")
	InOrderNotRecursion(tree)
	fmt.Println("\n ============ ")

	fmt.Println("postOrder")
	PostOrder(tree)
	fmt.Println("\n ============ ")

	fmt.Println("postOrderNotRecursion")
	PostOrderNotRecursion(tree)
	fmt.Println("\n ============ ")

	fmt.Println("BreadthFirstSearch")
	BreadthFirstSearch(tree)
	fmt.Println("\n ============ ")
}
