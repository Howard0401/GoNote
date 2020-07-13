package SingleLinkedList

import (
	"fmt"
)

//抽象方法
type Utility interface {
	GetFirstNode() *Node //抓取頭部節點

	InsertNodeFront(node *Node) //從頭部插入
	InsertNodeBack(node *Node)  //從尾部插入

	InsertTheDestNodeFront(dest *SingleList, node *Node) bool //從某個點前插入
	InsertTheDestNodeBack(dest *SingleList, node *Node) bool  //從某個點後插入

	GetNodeByIndex(index int) *Node
	DeleteNode(dest *Node) bool
	DeleteNodeByIndex(index int) //依index刪除
	// String() string              //顯示List
}

//實際上我們要操作的List
type SingleList struct {
	head   *Node
	length int
}

//建鍊表
func NewSingleList() *SingleList {
	head := NewNode(nil)
	// vi:=reflect.ValueOf(i)
	// return
	return &SingleList{head, 0}
}

//因為要對這個list操作
//Methods with pointer receivers(list *SingleList) can modify the value to which the receiver points
func (list *SingleList) GetFirstNode() *Node {
	return list.head.next
}

// 一開始是nil
// 如果不是nil就新建
func (list *SingleList) InsertNodeFront(node *Node) {

	if list.head == nil { //頭節點指向的位址為nil時
		list.head.next = node //頭節點為node
		node.next = nil       //指向的下一個節點是nil
		list.length++         //因為新插入節點，長度增加
	} else {
		// backup := list.head     //備份頭節點指向的位址{data,nil}
		// node.next = backup.next //把頭部節點指針指向的位址，賦值給node.next
		// backup.next = node      //再把原本的頭節點指向的位址改成node
		//跟上面是同等的意思，但這樣寫可以嗎？待確認
		node.next = list.head.next
		list.head.next = node
		list.length++ //因為新插入節點，長度增加
	}
}

func (list *SingleList) InsertNodeBack(node *Node) {
	if list.head == nil { //什麼都沒有就插入第一個節點
		list.head.next = node
		node.next = nil
		list.length++
	} else {
		bakeup := list.head
		for bakeup.next != nil { //如果發現節點不是最後一個，就往後遍歷到最後一個節點
			bakeup = bakeup.next
		}
		bakeup.next = node //插入
		list.length++
	}
}

func (list *SingleList) InsertTheDestNodeFront(dest interface{}, node *Node) bool {
	head := list.head
	isFound := false
	for head.next != nil { //如果找完這個鏈表前，能找到「某個點的下個節點」為要尋找的點，就標記
		if head.next.value == dest {
			isFound = true
			break
		}
		head = head.next
	}
	if isFound {
		node.next = head.next
		head.next = node
		list.length++
		return true
	} else {
		return false
	}
}

func (list *SingleList) InsertTheDestNodeBack(dest interface{}, node *Node) bool {
	head := list.head
	isFound := false
	for head.next != nil { //如果找完這鏈表前，能找到這個點，就標記現在head位置的的下個點
		if head.value == dest {
			isFound = true
			break
		}
		head = head.next
	}
	if isFound { //找到這個點後，
		node.next = head.next
		head.next = node
		list.length++
		return true
	} else {
		return false
	}
}

// https://golang.org/pkg/fmt/
//5. If an operand implements method String() string, that method will be invoked to convert the object to a string, which will then be formatted as required by the verb (if any).
func (list *SingleList) String() string {
	var result string
	p := list.head
	for p.next != nil {
		result += fmt.Sprintf("%v->", p.next.value)
		p = p.next
	}
	result += fmt.Sprintf("nil")
	return result
}

func (list *SingleList) GetNodeIndex(index int) *Node {
	if index > list.length-1 || index < 0 {
		fmt.Println("Please input correct index!")
		return nil
	} else {
		head := list.head
		for index > -1 { //向後找
			head = head.next
			index--
		}
		return head
	}
}

func (list *SingleList) DeleteNode(node *Node) bool {
	if node == nil {
		return false
	}
	head := list.head
	for head.next != nil && head.next != node {
		head = head.next
	}
	if head.next == node {
		head.next = head.next.next
		list.length--
		return true
	} else {
		return false
	}
}

func (list *SingleList) DeleteNodeByIndex(index int) {
	if index > list.length-1 || index < 0 {
		fmt.Println("Please input correct index!")
		return
	} else {
		head := list.head
		for index > 0 { //因為要找到前一個位置才能刪除
			head = head.next
			index--
		}
		head.next = head.next.next
		list.length--
		return
	}

}
