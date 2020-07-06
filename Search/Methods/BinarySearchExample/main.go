package main

import "fmt"

//根據業務需求作不同的變化

//重複的數，找該數的第一個索引
func BinSearchFirstIndex(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	index := -1

	for high >= low {
		mid := low + (high-low)/2 //避免溢位
		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else { //data == arr[mid]
			//return mid
			if mid == 0 || arr[mid-1] != data { //第一個數據 || 確保前面沒有一樣的數據
				index = mid
				fmt.Println("mid=", mid)
				break
			} else {
				high = mid - 1 //如果這個找到的數，前面還有一樣的數話，就遞迴繼續查找
			}
		}
	}
	return index
}

//重複的數，找該數的最後一個索引
func BinSearchLastIndex(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	index := -1

	for high >= low {
		mid := low + (high-low)/2 //避免溢位
		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else { //data == arr[mid]
			//return mid
			if mid == len(arr)-1 || arr[mid+1] != data { //最後一個數據 || 確保後面沒有重覆該數據
				index = mid
				fmt.Println("mid=", mid)
				break
			} else {
				low = mid + 1 //如果這個找到的數，後面還有一樣的數話，就遞迴繼續查找
			}
		}
	}
	return index
}

//找第一個>=data值的index
func BinSearchBiggerIndex(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	index := -1

	for high >= low {
		mid := low + (high-low)/2 //避免溢位
		if arr[mid] < data {
			low = mid + 1
		} else { //data == arr[mid]
			//return mid
			if mid == 0 || arr[mid-1] < data { //最後一個數據 || 確保後面沒有重覆該數據
				index = mid
				fmt.Println("mid=", mid)
				break
			} else {
				high = mid - 1 //如果這個找到的數，後面還有一樣的數話，就遞迴繼續查找
			}
		}
	}
	return index
}

//找第一個<=n值的index
func BinSearchSmallerIndex(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	index := -1

	for high >= low {
		mid := low + (high-low)/2 //避免溢位
		if arr[mid] > data {
			high = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] > data { //最後一個數據 || 確保後面沒有重覆該數據
				index = mid
				fmt.Println("mid=", mid)
				break
			} else {
				low = mid + 1 //如果這個找到的數，後面還有一樣的數話，就遞迴繼續查找
			}
		}
	}
	return index
}

func main() {
	arr := []int{1, 2, 3, 3, 3, 3, 3, 4, 5, 6, 6, 6, 6, 7, 9, 10}
	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[i], "index=", i)
	}
	fmt.Println(BinSearchFirstIndex(arr, 6))
	fmt.Println(BinSearchLastIndex(arr, 3))
	fmt.Println(BinSearchBiggerIndex(arr, 6))
	fmt.Println(BinSearchSmallerIndex(arr, 6))
}
