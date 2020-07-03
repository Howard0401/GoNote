package main

import "fmt"

/*
使用場合：高併發
1.取出依正整數n分割陣列空間的陣列，取頭尾兩端做比較交換
2.再將n-1繼續分割此陣列，並將頭尾比較交換，直到為0
//1 9 2 8     3 7 4 6 5 10  //以下為6格
//1             7
//	9             4
//	  2             6
//			8             5
//											10
//1 4 2 5     3 7 9 6 8 10 //再取5格依序往下排

/*
*/
/*
 */

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap {
		backup := arr[i]
		j := i - gap
		for j > 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
}

func ShellSort(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ { //處理每個元素的步長
				ShellSortStep(arr, i, gap)
			}
			// gap /= 2
			gap--
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(ShellSort(arr))
}
