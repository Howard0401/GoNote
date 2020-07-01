package main

import (
	"fmt"
	"strings"
)

//選取最大數
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

func SelectSortArr(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		for i := 0; i < len(arr)-1; i++ { //因為前面都是選出一個最大/最小的數加入排序後的arr，所以最後一個不用選
			min := i                            //標記索引
			for j := i + 1; j < len(arr); j++ { //選出一個最大值放進排序中
				if arr[min] < arr[j] {
					min = j
				}
			}
			if i != min { //如果選到最後一個數時，因為相等不用
				arr[i], arr[min] = arr[min], arr[i]
			} else {
				// fmt.Printf("%d", arr[i])
			}
			// fmt.Println(arr)
		}
	}
	return arr
}

func SelecSortMaxString(arr []string) string {
	if len(arr) <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] > max {
				// if strings.Compare(arr[i], max) > 0 { //這行可以取代
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSortArrString(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	} else {
		for i := 0; i < len(arr)-1; i++ { //因為前面都是選出一個最大/最小的數加入排序後的arr，所以最後一個不用選
			min := i                            //標記索引
			for j := i + 1; j < len(arr); j++ { //選出當前剩餘容器中的一個最大值，放進排序中
				// if arr[min] < arr[j] {
				if strings.Compare(arr[min], arr[j]) < 0 {
					min = j
				}
			}
			if i != min { //如果選到最後一個數時，因為相等不用
				arr[i], arr[min] = arr[min], arr[i]
			} else {
				// fmt.Printf("%d", arr[i])
			}
			// fmt.Println(arr)
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	arrString := []string{"c", "a", "b", "x", "z", "m", "n", "d", "f"}
	fmt.Printf("%d\n", SelecSortMax(arr))
	fmt.Println(SelectSortArr(arr))
	fmt.Printf("%v\n", SelecSortMaxString(arrString))
	fmt.Println(SelectSortArrString(arrString))
	/*
			//依序比較各個字母
			//left<right => -1
			fmt.Println(strings.Compare("b", "c"))
			//left= right => 0
			fmt.Println(strings.Compare("c", "c"))
			//left> right => 1
			fmt.Println(strings.Compare("c", "b"))


		pa := "a1"
		pb := "a2"
		fmt.Println("pa", &pa)
		fmt.Println("pb", &pb)
		fmt.Println(pa < pb) //go在1.1、1.3版本時比較的是address，go 1.10後可以這樣比較字串了
	*/
}
