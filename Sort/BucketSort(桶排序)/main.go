package main

import "fmt"

//當處理超大型數據、但數據範圍有限時，速度比快速排序還快
func BucketSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		//置入桶
		//num := 4 //已知要排序的數據種類數
		num := length
		max := SelecSortMax(arr)
		index := 0 //索引
		buckets := make([][]int, num)
		for i := 0; i < length; i++ {
			index = arr[i] * (num - 1) / max //桶自動分配(會有空桶)
			buckets[index] = append(buckets[index], arr[i])
			//buckets[arr[i]-1] = append(buckets[arr[i]-1], arr[i]) //桶數+1
			//為何要這樣寫呢？ 這樣buckets才可以從[0]開始
		}
		fmt.Println(buckets)
		//排序
		tmp := 0
		for i := 0; i < num; i++ {
			bucketslen := len(buckets[i]) //求某段的長度
			if bucketslen > 0 {           //如果長度>0
				buckets[i] = SelectSortArr(buckets[i]) //桶內數據排序(為啥需要這個)
				copy(arr[tmp:], buckets[i])            //copy該桶的數據到arr的起始段
				tmp += bucketslen                      //再追加長度
			}
		}
		return arr
	}
}

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

func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 2, 3, 1} //可分成 11 2222 333 44這四種
	fmt.Println(BucketSort(arr))
}
