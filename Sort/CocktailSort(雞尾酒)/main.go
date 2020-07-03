package main

import "fmt"

/*
雙向的BuubleSort(其實這句就夠了)
(下面解釋未必正確 之後檢查再修改)


// 8 1 4 2 9 5 3

// 1 4 2 8 (前面8依序比到9時，因為比9還小，所以固定8的位置，接著從8往後繼續比較) 5 3 9

	 1 4  2 8 5 3 9

***接下來反向操作

   1 2 4(因為9排好了，選擇9的前一個值3，當3交換到前面，發現有比它更小的值後，固定3當下的位置，將更小的值繼續往前換)3 8 5 9

*/

func CocktailSort(arr []int) []int {
	length:= len(arr)
	for i:=0; i<length/2; i++{
		left:=0
		right:=length-1
		for left<=right{//結束的條件
			if arr[left] > arr[left+1]{
				arr[left],arr[left+1]= arr[left+1],arr[left]
			}
			left++
			if arr[right-1]>arr[right]{
				arr[right-1],arr[right] = arr[right],arr[right-1]
			}
			right--
		}
		fmt.Println(i,arr)
	}
	return arr
}


func main() {
	arr:=[]int {1,9,2,8,3,7,4,6,5,10}
	fmt.Println(CocktailSort(arr))
}