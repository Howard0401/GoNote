package main

import "fmt"

//Better Shell 希爾排序改良版

func ComoSort(arr []int) []int {
	length := len(arr)
	gap := length
	for gap > 1 {
		gap = gap * 10 / 13
		for i := 0; i+gap < length; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(ComoSort(arr))
}
