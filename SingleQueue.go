//Front，值為放進去的flag，預設為-1
//Rear，值為放輸入新數值的flag，預設為-1  Maxsize-1為Queue的總長度

//addqueue
package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	slice   []int //改成Slice了 //陣列(數組
	front   int   //queue flag
	rear    int   //input flag
}

//添加數據到Queue的架構
func (this *Queue) Add(val int) (err error) {
	//先判Queue是否已滿
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.rear++ //如果添加成功
	this.slice[this.rear] = val
	return
}

func (this *Queue) GetQueue() (val int, err error) {
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.slice[this.front]
	return val, err
}

func (this *Queue) ShowQueue() {
	//this.front不包含隊首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("slice[%d]=%d\n", i, this.slice[i])
	}
}

func main() {
	//Initialize
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
		slice:   make([]int, 5),
	}

	var key string
	var val int

	for {
		fmt.Println("1.輸入add表示添加數據到Queue")
		fmt.Println("2.輸入get表示從Queue獲取數據")
		fmt.Println("3.輸入show顯示Queue")
		fmt.Println("4.結束並顯示Queue")

		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Println("輸入添加進Queue的值")
			fmt.Scanln(&val)
			err := queue.Add(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("取出的數為:", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
