package main

import "fmt"

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		//Ver1:想像有兩條arr在互相比較
		// for i := 0; i < len(arr); i++ {
		// 	index := false //
		// 	for j := i + 1; j < len(arr); j++ {
		// 		if arr[j] < arr[i] {
		// 			arr[i], arr[j] = arr[j], arr[i]
		// 			index = true
		// 		}
		// // 	}
		//Ver2:同一個陣列每次比較往後移動n格(j) 比較n-1次(i)
		for i := 0; i < len(arr)-1; i++ {
			index := false //
			for j := 0; j < len(arr)-1-i; j++ {
				if arr[j+1] < arr[j] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					index = true
				}
			}
			if !index {
				break
			}
		}
		return arr
	}
}

func main() {
	arr := []int{3, 2, 4, 5, 2, 6, 1}
	fmt.Println(BubbleSort(arr))
}
