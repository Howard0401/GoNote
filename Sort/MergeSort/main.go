package main

import "fmt"

func Merge(leftArr []int, rightArr []int) []int {
	leftIndex := 0
	rightIndex := 0
	lastArr := []int{} //最終回傳的值
	for leftIndex < len(leftArr) && rightIndex < len(rightArr) {
		//第一輪是把兩個僅有一個值{}的Arr做排序 大的排前面小的排後面
		//如{7,1}經排序後會變成{1,7}，這樣排完後第一輪就結束了
		//第二輪假設我們上輪解出的值是{1,7}作為left 而right則為{3,9}
		//依照這樣排序的邏輯，可以排出{1,3,7,9}這樣的Arr。
		if leftArr[leftIndex] < rightArr[rightIndex] {
			lastArr = append(lastArr, leftArr[leftIndex])
			leftIndex++

		} else if leftArr[leftIndex] > rightArr[rightIndex] {
			lastArr = append(lastArr, rightArr[rightIndex])
			rightIndex++

		} else {
			lastArr = append(lastArr, leftArr[leftIndex])
			lastArr = append(lastArr, rightArr[rightIndex])
			leftIndex++
			rightIndex++
		}
	}

	//這種情況是因為當初分開left和right後，排序合併時左或右有可能不對稱遺留一個單獨的數
	//因此必須把這個數單獨拿出來看要插入在
	for leftIndex < len(leftArr) { //把沒有結束的Merge進去
		lastArr = append(lastArr, leftArr[leftIndex])
		leftIndex++
	}

	for rightIndex < len(rightArr) {
		lastArr = append(lastArr, rightArr[rightIndex])
		rightIndex++
	}
	return lastArr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		mid := length / 2
		left := MergeSort(arr[:mid])
		right := MergeSort(arr[mid:])
		return Merge(left, right)
	}

}

func main() {
	arr := []int{9, 2, 1, 6, 0, 7}
	fmt.Println(MergeSort(arr))
}
