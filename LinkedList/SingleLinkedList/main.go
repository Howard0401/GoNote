package main

import (
	SingleLinkedList "SingleLinkedList/Packages"
	"fmt"
)

func main() {
	list := SingleLinkedList.NewSingleList()
	fmt.Println(list)
	node1 := SingleLinkedList.NewNode(1)
	node2 := SingleLinkedList.NewNode(2)
	node3 := SingleLinkedList.NewNode(3)
	node4 := SingleLinkedList.NewNode(4)
	node5 := SingleLinkedList.NewNode(5)
	list.InsertNodeFront(node1)
	fmt.Println(list)
	list.InsertNodeBack(node2)
	fmt.Println(list)
	list.InsertNodeBack(node3)
	fmt.Println(list)

	list.InsertTheDestNodeBack(2, node4)
	fmt.Println(list)
	list.InsertTheDestNodeFront(2, node5)
	fmt.Println(list)
	fmt.Println(list.GetNodeIndex(2))
	list.DeleteNode(node3) //刪除一個節點
	fmt.Println(list)
	list.DeleteNodeByIndex(2)
	fmt.Println(list)
}
