package main

import "fmt"

type Node struct {
	value int
	next *Node
}

type LinkedList struct { 
	head *Node
}
func (list *LinkedList) add(value int) {
    newNode := &Node{value: value}

    if list.head == nil {
        list.head = newNode
        return
    }

	current := list.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

func (list *LinkedList) Print() {
    current := list.head
    for current != nil {
        fmt.Print(current.value, " ")
        current = current.next
    }
    fmt.Println()
}

func main() {
	list := &LinkedList{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	list.add(5)
 
    list.Print() 
}