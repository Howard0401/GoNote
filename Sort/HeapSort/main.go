package main

import "fmt"

/*
1.選出Binary Heap中最大的數值
2.再把當前最大的數放到後面，重複進行獲得
*/

func HeapSortMax(arr []int, length int) []int {
	// length := len(arr)
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1         //深度 n 2*n+1 2*n+2
		for i := depth; i >= 0; i-- { //循環所有的n節點(這邊每個n代表由top left right節點構成的架構)
			top := i
			left := 2*i + 1
			right := 2*i + 2
			if arr[left] > arr[top] && left <= length-1 { //如果左邊比頂部大，且左節點存在，就記錄在頂部
				top = left
			}
			if arr[right] > arr[top] && right <= length-1 { //如果右邊比頂部大，且右節點存在，替換頂部
				top = right
			}
			if top != i {
				arr[i], arr[top] = arr[top], arr[i]
			}
		}
		return arr
	}
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		last := length - i
		HeapSortMax(arr, last)
		if i < length {
			arr[0], arr[last-1] = arr[last-1], arr[0]
		}
	}
	return arr
}

func main() {
	arr := []int{3, 2, 4, 5, 2, 6, 1}
	fmt.Println(HeapSort(arr))
}
