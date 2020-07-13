// package main123

// import "fmt"

// //Data Structure of List
// type HeroNode struct {
// 	no       int //Data
// 	name     string
// 	nickname string
// 	next     *HeroNode //Pointer to next
// }

// //Insert (無法依序)
// func InsertNode(head *HeroNode, newNode *HeroNode) {
// 	//1.先找到List的最後節點

// 	//2.創建一個輔助節點
// 	temp := head
// 	for {
// 		if temp.next == nil {
// 			break
// 		}
// 		temp = temp.next //上一個非空就指向下個要新增的next
// 	}
// 	//3.將Node加入到最尾
// 	temp.next = newNode
// }

// //Insert(依編號排序)
// func InsertNode2(head *HeroNode, newNode *HeroNode) {
// 	temp := head
// 	flag := true
// 	for {
// 		if temp.next == nil {
// 			break
// 		} else if temp.next.no > newNode.no { //大於是指插入的資料要在原先資料的後面,所以放到下面處理
// 			break
// 		} else if temp.next.no == newNode.no {
// 			flag = false
// 			break
// 		}
// 		temp = temp.next
// 	}

// 	if !flag {
// 		fmt.Println("The Node is Exist!!:", newNode.no)
// 		return
// 	} else {
// 		newNode.next = temp.next
// 		temp.next = newNode
// 	}

// }

// //Delete
// func Delete(head *HeroNode, id int) {
// 	temp := head
// 	flag := false
// 	for {
// 		if temp.next == nil {
// 			break
// 		} else if temp.next.no == id {
// 			flag = true
// 			break
// 		}
// 		temp = temp.next
// 	}
// 	if flag {
// 		temp.next = temp.next.next
// 	} else {
// 		fmt.Println("id不存在")
// 	}
// }

// //Display
// func ListNode(head *HeroNode) {
// 	temp := head
// 	//1.是不是空的
// 	if temp.next == nil {
// 		fmt.Println("List is Empty")
// 		return
// 	}
// 	//2.非空至少有一個節點就印出來
// 	for {
// 		if temp.next.next != nil { //發現指向下個節點的指標為nil後，就不輸出箭頭 ps 因為第一個是head，第二個是這筆資料，所以才寫temp.next.next
// 			fmt.Printf("[%d,%s,%s]===>", temp.next.no, temp.next.name, temp.next.nickname)
// 			temp = temp.next
// 		} else {
// 			fmt.Printf("[%d,%s,%s]", temp.next.no, temp.next.name, temp.next.nickname)
// 			break
// 		}
// 	}
// }

// func main() {
// 	//需要一個空的Head給後面串
// 	head := &HeroNode{}

// 	//建立下個節點
// 	hero1 := &HeroNode{
// 		no:       1,
// 		name:     "第一筆",
// 		nickname: "第一筆資料",
// 	}

// 	hero2 := &HeroNode{
// 		no:       2,
// 		name:     "第二筆",
// 		nickname: "第二筆資料",
// 	}

// 	hero3 := &HeroNode{
// 		no:       3,
// 		name:     "第三筆",
// 		nickname: "第三筆資料",
// 	}

// 	// InsertNode(head, hero1)
// 	// InsertNode(head, hero2)
// 	// InsertNode(head, hero3)
// 	InsertNode2(head, hero2)
// 	InsertNode2(head, hero3)
// 	InsertNode2(head, hero1)

// 	ListNode(head)
// 	fmt.Println()
// 	Delete(head, 2)
// 	ListNode(head)

// }
