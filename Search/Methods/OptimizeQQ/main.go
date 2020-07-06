package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type QQ struct {
	QQuser   int
	password string
}

func swap(arr []QQ, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSortX(arr []QQ, left, right int) { //遞迴左小右大
	if right-left < 15 { //如遞迴時陣列間距<3，插入排序
		InsertionSort(arr, left, right)
	} else {
		swap(arr, left, rand.Int()%(right-left+1)+left) //隨機找一個數放在第一位
		vData := arr[left]
		lt := left      //要使arr[left+1...lt] < vData
		gt := right + 1 //要使arr[gt...right] > vData
		i := left + 1   //要使arr[lt+1...i] == vData
		for i < gt {
			if arr[i].QQuser < vData.QQuser { //往左
				swap(arr, i, lt+1) //移動到小於的地方
				lt++
				i++
			} else if arr[i].QQuser > vData.QQuser { //往右
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

func InsertionSort(arr []QQ, left, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i]
		var j int
		for j = i; j > left && arr[j-1].QQuser > temp.QQuser; j-- { //定位
			arr[j] = arr[j-1] //選取插入位後，原本的數向後移
		}
		arr[j] = temp //插入
	}
}

func QuickSortPlus(arr []QQ) { //快速排序的核心
	QuickSortX(arr, 0, len(arr)-1)
}

func BinSearch(arr []QQ, data int) int {
	low := 0
	high := len(arr) - 1
	for high >= low {
		mid := (high + low) / 2
		if arr[mid].QQuser > data {
			high = mid - 1
		} else if arr[mid].QQuser < data {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	path := "..\\..\\QQBig.txt"
	file, _ := os.Open(path)
	br := bufio.NewReader(file)

	//Assign DataN
	allData := make([]QQ, 0)
	i := 0
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		lineStr := string(line)
		lines := strings.Split(lineStr, "----")
		user, err := strconv.Atoi(lines[0])
		if len(lines) == 2 {
			in := QQ{
				QQuser:   user,
				password: lines[1],
			}
			allData = append(allData, in)
		} else {
			fmt.Println(err)
		}
		i++
	}
	fmt.Println("載入完成")
	starttime := time.Now()
	QuickSortPlus(allData)
	fmt.Println("排序完成")
	fmt.Println("本次排序用了", time.Since(starttime))
	flag := true
	for flag {
		fmt.Println("請輸入欲查詢的用戶名") //ex.58243449
		var inputstr int
		fmt.Scanln(&inputstr)                 //用户输入
		index := BinSearch(allData, inputstr) //用Binary取代linear

		if index == -1 {
			fmt.Println("找不到")
		} else {
			fmt.Println("找到", allData[index])
		}
		fmt.Println("是否繼續查詢?(Y/N)")
		key := ""
		fmt.Scanln(&key)
		if key == "Y" {
			flag = true
		} else {
			break
		}

	}
}
