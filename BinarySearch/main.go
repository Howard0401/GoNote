package main

import "fmt"

//Ver1
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

//Ver2
func SortForMerge(arr []int, left int, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i]
		var j int
		for j :=i; j>left &&
	}
}

func InserTest(arr[] int)[]int{
	j:=
}

//Ver2 recursion
func QuickSortX(arr []int, left int, right int) []int {
	if right-left < 3 { //Slice中剩下3格字時 直接插入排序

	} else {

	}
}

//Ver2
func QuickSortPlus(arr []int) {
	QuickSortX()
}

func bin_search(arr []int, data int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2
		if data < arr[mid] {
			right = mid - 1
		} else if data > arr[mid] {
			left = mid + 1
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
	fmt.Println(bin_search(arr, 4))
	index := bin_search(arr, 19)
	if index == -1 {
		fmt.Println("沒有找到")
	} else {
		fmt.Println("找到", arr[index], index)
	}
}
