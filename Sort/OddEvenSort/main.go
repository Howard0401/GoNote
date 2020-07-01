package main

import "fmt"

//可以多執行緒
/*
從奇數位的arr做BubbleSot一次，再從arr的偶數位排序一次，持續迭代
9 2 1 6 0 7  //Origin
 //odd
2 9 1 6 0 7
//even
2 1 9 0 6 7
//odd
1 2 0 9 6 7
//even
1 2 9 0 6 7
//odd
1 2 0 9 6 7
//even
1 0 2 6 9 7
//odd
0 1 2 6 7 9
應用的場合；身高的方差
*/

func OddEvenSort(arr []int) []int {
	lenth := len(arr)
	isSorted := false //如果沒有更動過排序時，跳出迴圈的flag
	for !isSorted {
		isSorted = true
		for i := 1; i < lenth-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		for i := 0; i < lenth-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr
}

func main() {
	arr := []int{9, 2, 1, 6, 0, 7}
	fmt.Println(arr)
	fmt.Println(OddEvenSort(arr))

}
