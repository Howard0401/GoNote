package main

import (
	"fmt"
	"math/rand"
)

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSortX(arr []int, left, right int) { //遞迴左小右大
	if right-left < 2 { //如遞迴時陣列間距<3，插入排序
		InsertionSort(arr, left, right)
	} else {
		swap(arr, left, rand.Int()%(right-left+1)+left) //隨機找一個數放在第一位
		vData := arr[left]
		lt := left      //要使arr[left+1...lt] < vData
		gt := right + 1 //要使arr[gt...right] > vData
		i := left + 1   //要使arr[lt+1...i] == vData
		for i < gt {
			if arr[i] < vData { //往左
				swap(arr, i, lt+1) //移動到小於的地方
				lt++
				i++
			} else if arr[i] > vData { //往右
				swap(arr, i, gt-1) //移動到大於的地方
				gt--
			} else { //等於
				i++
			}
		}
		swap(arr, left, lt) //交換頭部的位置
		//遞迴處理各段大於、小於的部分
		QuickSortX(arr, left, lt-1)
		QuickSortX(arr, gt, right)
	}
}

func InsertionSort(arr []int, left, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i]
		var j int
		for j = i; j > left && arr[j-1] > temp; j-- { //定位
			arr[j] = arr[j-1] //選取插入位後，原本的數向後移
		}
		arr[j] = temp //插入
	}

}

func QuickSortPlus(arr []int) { //快速排序的核心
	QuickSortX(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println("排序前", arr)
	QuickSortPlus(arr)
	fmt.Println("排序後", arr)
}
