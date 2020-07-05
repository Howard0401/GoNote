package main

import "fmt"

func MidSearch(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	i := 0
	for high >= low {
		i++
		fmt.Println("這是第幾次?", i)
		Resize := float64(data - arr[low])           //--- 3
		leftToRight := float64(arr[high] - arr[low]) //----- 5 比例就是3/5
		Originlength := float64(high - low)
		mid := int(float64(low) + Originlength*Resize/leftToRight)
		if mid < 0 || mid > len(arr) { //找不到越界時
			return -1
		}
		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println("index=", MidSearch(arr, 124))
}
