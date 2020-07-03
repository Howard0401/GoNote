package main

import "fmt"

func SelecSortMax(arr []int) int {
	if len(arr) <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

func RadixSort(arr []int) []int {
	max := SelecSortMax(arr)                //尋找極大值 999 99991
	for bit := 1; max/bit > 10; bit *= 10 { //按數量級分段
		arr = BitSort(arr, bit) //先處理個位數，再處理十位和百位
		fmt.Println(arr)
	}
	return arr
}

func BitSort(arr []int, bit int) []int {
	length := len(arr)
	bitCounts := make([]int, 10)
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10 //分層，當bit=1000的時候 三位數就不再排序了，而當bit=10000時，四位數就不再排序了
		bitCounts[num]++           //統計餘數相等個數
	}
	fmt.Println(bitCounts)
	// 0 1 2 3 4 5  -(1)
	// 1 0 3 0 0 1  -(2)
	// 1 1 4 4 4 5
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1]
	}
	fmt.Println(bitCounts)
	tmp := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitCounts[num]-1] = arr[i]
		bitCounts[num]--
	}
	for i := 0; i < length; i++ {
		arr[i] = tmp[i]
	}
	return arr
}

func main() {
	arr := []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 1011, 1111022}
	fmt.Println(RadixSort(arr))

}
