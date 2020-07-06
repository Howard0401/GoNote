package main

import "fmt"

//This is mid Search
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

//FibonacciSearch
//拆分方式依照Fibonachii
func FiboSearch(arr []int, val int) int {
	length := len(arr)
	fabArr := FiboArray(arr) //製造Fibonachii數列
	// fmt.Println(fabArr)
	LengthToFill := fabArr[len(fabArr)-1]
	fmt.Println(LengthToFill)
	ArrayToFill := make([]int, LengthToFill)
	// for i, v := range arr {
	// 	ArrayToFill[i] = v
	// }//下面是這個的簡易寫法
	copy(ArrayToFill, arr)
	lastData := arr[length-1] //填充最後一個大數

	for i := length; i < LengthToFill; i++ {
		ArrayToFill[i] = lastData //填充以後的數
	}
	// fmt.Println(ArrayToFill, length, len(ArrayToFill))

	left, mid, right := 0, 0, length
	kindex := len(fabArr) - 1

	for left <= right {
		mid = left + fabArr[kindex-1] - 1 //使用Fabonachi切割
		if val < ArrayToFill[mid] {
			right = mid - 1
			kindex--
		} else if val > ArrayToFill[mid] {
			left = mid + 1
			kindex -= 2 //反覆循環兩次?
		} else {
			if mid > right {
				return right //越界
			} else {
				return mid
			}
		}
	}
	return -1
}

func FiboArray(arr []int) []int {
	length := len(arr)
	// fmt.Println(length)
	fiblen := 2
	first := 1
	second, third := 1, 2
	for third < length && first >= 0 { //找出最近的fibo
		first, second, third = second, third, second+third //疊加數據
		// fmt.Println(third, first, second)
		fiblen++ //fibonachii的下標
	}
	fmt.Println(fiblen)
	fb := make([]int, fiblen)
	fb[0] = 1
	fb[1] = 1
	for i := 2; i < fiblen; i++ {
		fb[i] = fb[i-1] + fb[i-2] //疊加計算
	}
	return fb
}

func main() {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i * 100
	}
	// fmt.Println("index=", MidSearch(arr, 124))
	// fmt.Println(FiboArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	// fmt.Println(FiboArray(arr))
	index := FiboSearch(arr, 233)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("找不到")
	} else {
		fmt.Println("找到:", index)
	}
}
