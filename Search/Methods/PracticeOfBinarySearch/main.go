package main

import "fmt"

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 { //只有一個
		return arr
	} else {
		splitData := arr[0] //第一個數
		lower := make([]int, 0)
		higher := make([]int, 0)
		mid := make([]int, 0)
		mid = append(mid, splitData) //保存分離的數
		for i := 1; i < length; i++ {
			if arr[i] < splitData {
				lower = append(lower, arr[i])
			} else if arr[i] > splitData {
				higher = append(higher, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		lower, higher = QuickSort(lower), QuickSort(higher)
		myarr := append(append(lower, mid...), higher...)
		return myarr
	}
}

func BinSearch(arr []int, data int) int {
	low := 0
	hight := len(arr) - 1
	for hight >= low { //等於下個回合就是比較切一半的值
		// fmt.Println(arr[low:hight])
		mid := (low + hight) / 2 //取中間值
		fmt.Println(mid)
		if arr[mid] > data { //移到前面的分割
			hight = mid - 1
		} else if arr[mid] < data { //移到後面的分割
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 19, 4, 8, 3, 5, 4, 6, 19, 0}
	fmt.Println("未經排序", arr)
	fmt.Println("經排序後", QuickSort(arr))
	arr = QuickSort(arr)
	index := BinSearch(arr, 3)
	if index == -1 {
		fmt.Println("找不到")
	} else {
		fmt.Println("找到index=", index)
	}
}
