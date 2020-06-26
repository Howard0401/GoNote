/*
1.Initial: tail=0 head=0

2.如何判斷Queue滿?
(tail + 1) % MaxLenth == head (為何如此做的理由寫在5.)

3.判 斷為空
tail==head

4.怎麼統計Queue中有多少個元素
(tail - head + MaxLenth) % MaxLenth

5.為什麼Queue要留最後一格空間呢?
可以觀察到上面判斷已滿和有幾個元素的方法，都使用MaxLenth來檢查

*/

package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	MaxLenth int
	slice    []int
	head     int //頭 Init 0
	tail     int //尾 Init 0
}

//Full
func (this *CircleQueue) Isfull() bool {
	return (this.tail+1)%this.MaxLenth == this.head
}

//Empty
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head

}

//Size
func (this *CircleQueue) Size() int {
	return (this.tail - this.head + this.MaxLenth) % this.MaxLenth
}

//List
func (this *CircleQueue) ListQueue() {
	fmt.Println("Queue Contains：")
	size := this.Size()
	if size == 0 {
		fmt.Println("Queue is Empty")
	}
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.slice[tempHead])
		tempHead = (tempHead + 1) % this.MaxLenth
	}
}

//Push
func (this *CircleQueue) Push(val int) (err error) {
	if this.Isfull() {
		return errors.New("The queue is FULL!!")
	}
	//this.tail
	this.slice[this.tail] = val
	this.tail = (this.tail + 1) % this.MaxLenth
	return
}

//Pop
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.slice[this.head]
	this.head = (this.tail + 1) % this.MaxLenth
	return
}

func main() {
	//Initialize
	queue := &CircleQueue{
		MaxLenth: 5,
		slice:    make([]int, 5),
		head:     0,
		tail:     0,
	}

	var key string
	var val int

	for {
		fmt.Println("1.輸入push表示添加數據到Queue")
		fmt.Println("2.輸入pop表示從Queue獲取數據")
		fmt.Println("3.輸入list顯示Queue")
		fmt.Println("4.輸入exit並結束Queue")

		fmt.Scanln(&key)

		switch key {
		case "push":
			fmt.Println("輸入添加進Queue的值")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入成功")
			}
		case "pop":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("取出的數為:", val)
			}
		case "list":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
