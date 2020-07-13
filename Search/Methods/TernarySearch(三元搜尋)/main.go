package main

import "fmt"

//為什麼要學三分查找？
//分布式在不同機器上時可使用

func TernarySearch(arr []int, data int) int { //簡單說就是把標籤變成兩個、三份
	low := 0
	high := len(arr) - 1
	i := 0
	for high >= low {
		i++
		mid1 := low + int((high-low)/3)
		mid2 := high - int((high-low)/3)
		midData1 := arr[mid1]
		midData2 := arr[mid2]
		if midData1 == data {
			return mid1
		} else if midData2 == data {
			return mid2
		}

		if midData1 < data {
			low = mid1 + 1
		} else if midData2 > data {
			high = mid2 - 1
		} else {
			low = low - 1
			high = high - 1
		}
		fmt.Println("次數", i)
	}

	return -1
}

func main() {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println(TernarySearch(arr, 12))
}
