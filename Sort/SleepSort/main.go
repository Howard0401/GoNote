package main

import (
	"fmt"
	"time"
)

//依照各個數值的大小，分別休眠相異時間後再排序
//多執行續(Tread)、分布式都很適合
//那如果是字串要怎麼處理？=>使用Hash

var flag bool
var container chan bool
var count int

func main() {
	var arr []int = []int{16, 8, 1, 24, 30}
	flag = true
	container = make(chan bool, 5) //開闢5個channel
	for i := 0; i < len(arr); i++ {
		go ToSleep(arr[i])
	}
	go Listen(len(arr))
	for flag {
		time.Sleep(time.Nanosecond)
	}
}

func ToSleep(data int) {
	time.Sleep(time.Duration(data) * time.Microsecond * 100)
	fmt.Println("Sleep", data)
	container <- true
}

func Listen(size int) {
	for flag { //輸入為真，處理channel
		select {
		case <-container:
			count++
			if count >= size {
				flag = false
				break
			}
		}
	}
}
