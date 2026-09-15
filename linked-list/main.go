package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (l *LinkedList) String() string {
	result := ""

	for current := l.Head; current != nil; current = current.Next {
		result += fmt.Sprintf("%d -> ", current.Value)
	}
	return result + "nil"
}

func main() {
	list := &LinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	fmt.Println(list)
}
