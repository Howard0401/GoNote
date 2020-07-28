package main

import (
	doublyLinkedList "DoubleLinkedList/Packages"
	"fmt"
)

func main() {
	list := doublyLinkedList.NewList()
	node1 := doublyLinkedList.NewNode(1)
	node2 := doublyLinkedList.NewNode(2)
	node3 := doublyLinkedList.NewNode(3)
	node4 := doublyLinkedList.NewNode(4)
	node5 := doublyLinkedList.NewNode(5)
	list.InsertFront(node1)
	list.InsertFront(node2)
	list.InsertFront(node3)
	list.InsertBack(node4)
	list.InsertBack(node5)
	fmt.Println(list)
	// fmt.Println(list.String())
}
