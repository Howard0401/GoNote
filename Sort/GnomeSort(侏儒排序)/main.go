package main

import "fmt"

//用途；假設快速排序很多個(1億個)，在分割小片段排序時可用侏儒排序
func GnomeSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		if arr[i] >= arr[i-1] { //如果由前往後、從小排到大，符合條件就繼續往後找
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 { //如果交換成功，就把被交換的那個數和比較的index往前挪
				i--
			}
		}

	}
	return arr
}

func main() {
	arr := []int{11, 2, 3, 23, 33, 3, 13, 4, 15, 6, 6, 61, 6, 17, 9, 10}
	fmt.Println(GnomeSort(arr))
}
