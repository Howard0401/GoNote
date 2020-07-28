package main

import (
	SingleLinkedList "SingleLinkedList/Packages"
	"fmt"
)

//什麼時候用LinkedList比較好？需要頻繁的刪除和插入，但不常查詢(像工作管理員是Doubly LinkedList)

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
	// fmt.Println(list.GetMid().Value())
	list.ReverseList()
	fmt.Println(list)
	list.ReverseList()
	fmt.Println(list)
}

// func main() {
// 	list := SingleLinkedList.NewSingleList()
// 	path := "./1_3.txt"
// 	// path := "./CreditCardData.txt"
// 	file, _ := os.Open(path)
// 	br := bufio.NewReader(file)
// 	i := 0
// 	for {
// 		line, _, end := br.ReadLine()
// 		if end == io.EOF {
// 			break
// 		}
// 		lineStr := string(line)
// 		nodeStr := SingleLinkedList.NewNode(lineStr)
// 		list.InsertNodeFront(nodeStr)
// 		i++
// 	}
// 	fmt.Println(i, "存取記憶體完成")
// 	flag := true
// 	for flag {
// 		fmt.Println("請輸入要查詢的用戶名")
// 		var inputStr string
// 		fmt.Scanln(&inputStr)
// 		fmt.Println(inputStr)
// 		startTime := time.Now()
// 		list.FindString(inputStr)
// 		fmt.Println("本次查询用了", time.Since(startTime))

// 		fmt.Println("是否繼續查詢?(Y/N)")
// 		key := ""
// 		fmt.Scanln(&key)
// 		if key == "Y" {
// 			flag = true
// 		} else {
// 			break
// 		}
// 	}
// }
