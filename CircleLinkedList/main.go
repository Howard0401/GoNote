/*

環形列表head放Dta

*/
package main

import "fmt"

type Node struct {
	no   int
	name string
	next *Node
}

func Insert(head *Node, newNode *Node) {

	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head //構成環形
		return           //已加入第一筆Datas
	}

	temp := head //為了找到環形的最後節點
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newNode //先建立一條新的線，指向要插入的節點
	newNode.next = head //The pointer to head equals the newNode.next poiter to head

}

/*
1.先讓temp指向head
2.定義helper指向環形LinkedList的最後
3.讓temp和刪除的id比較，若相同通過helper刪除(但如果刪的是第一個?)

這樣講可能比較清楚：
定義一個在當前temp之前的標籤help，比對序號時若temp==id
就把help指向刪除後該節點後的下個節點

*/
func Delete(head *Node, id int) *Node {

	temp := head
	helper := head

	if temp.next == nil {
		fmt.Println("LinkedList is Empy!!")
		return head
	}

	if temp.next == head && temp.no == id { //只有一個節點
		temp.next = nil
		return head
	}

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	//如果有兩個以上的節點
	for {
		if temp.next == head { //已經比較到最後一個，但最後一個還沒找過
			break
		}
		if temp.no == id { //找到
			if temp == head { //如果刪除第一個節點就需要修改head的下個節點
				head = head.next
			}
			helper.next = temp.next //刪除 這邊helper.next是
			fmt.Printf("刪除%d節點\n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}

	//最後一個
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("刪除%d節點\n", id)
		} else {
			fmt.Printf("沒有這個id:%d\n", id)
		}
	}
	return head
}

func ShowList(head *Node) {
	temp := head
	if temp.next == nil {
		fmt.Println("Empty List")
		return
	}
	for {
		fmt.Printf("[%d,%s]=>", temp.no, temp.name)
		if temp.next == head {
			fmt.Println()
			break
		}
		temp = temp.next
	}
}

func main() {
	//head要有值
	head := &Node{}
	node1 := &Node{
		no:   1,
		name: "第一筆",
	}
	node2 := &Node{
		no:   2,
		name: "第二筆",
	}
	node3 := &Node{
		no:   3,
		name: "第三筆",
	}

	Insert(head, node1)
	Insert(head, node2)
	Insert(head, node3)
	ShowList(head)
	fmt.Println()
	head = Delete(head, 1)

	head = Delete(head, 2)
	head = Delete(head, 1)
	ShowList(head)
	head = Delete(head, 3)
	head = Delete(head, 3)
	ShowList(head)
}
