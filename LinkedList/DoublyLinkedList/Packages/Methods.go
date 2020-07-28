package doublyLinkedList

import "fmt"

type DoublyList struct {
	head   *Node
	length int
}

//Create
func NewList() *DoublyList {
	head := NewNode(nil)
	return &DoublyList{head, 0}
}

func (list *DoublyList) GetLength() int {
	return list.length
}

func (list *DoublyList) GetFirstNode() *Node {
	return list.head.next
}

func (list *DoublyList) InsertFront(node *Node) {
	head := list.head
	if head.next == nil {
		node.next = head.next
		head.next = node
		node.prev = head
		list.length++
	} else {
		head.next.prev = node
		node.next = head.next
		head.next = node
		node.prev = head
		list.length++
	}
}

func (list *DoublyList) InsertBack(node *Node) {
	head := list.head
	if head.next == nil {
		node.next = head.next
		head.next = node
		node.prev = head
		list.length++
	} else {
		for head.next != nil {
			head = head.next
		}
		head.next = node
		node.prev = head
		list.length++
	}
}

func (list *DoublyList) String() string {
	var listString1 string
	var listString2 string
	head := list.head
	//從左到右
	listString1 += fmt.Sprintf("nil-->")
	for head.next != nil {
		listString1 += fmt.Sprintf("%v-->", head.next.value)
		head = head.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"

	//從右到左
	listString1 += fmt.Sprintf("nil-->")
	for head != list.head {
		listString2 += fmt.Sprintf("%v-->", head.value)
		head = head.prev
	}
	listString2 += fmt.Sprintf("nil")
	listString2 += "\n"

	return listString1 + listString2 + "\n"
}
