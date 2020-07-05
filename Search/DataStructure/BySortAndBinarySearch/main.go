package main

/*
1.建立包含使用者ID與密碼的資料結構
2.逐行讀入並存取在allStr的slice中
3.先經由快速排序，再進行Binary Search
4.找出符合的ID後，退出或再執行
*/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type QQ struct {
	QQuser int
	QQpass string
}

const N = 84331469

func QuickSortStruct(arr []QQ) []QQ {
	length := len(arr)
	if length <= 1 { //只有一個
		return arr
	} else {
		splitData := arr[0].QQuser //第一個數
		lower := make([]QQ, 0)
		higher := make([]QQ, 0)
		mid := make([]QQ, 0)
		// mid = append(mid, splitData) //保存分離的數 為啥這邊要這樣改?
		mid = append(mid, arr[0])
		for i := 1; i < length; i++ {
			if arr[i].QQuser < splitData {
				lower = append(lower, arr[i])
			} else if arr[i].QQuser > splitData {
				higher = append(higher, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		lower, higher = QuickSortStruct(lower), QuickSortStruct(higher)
		myarr := append(append(lower, mid...), higher...)
		return myarr
	}
}

func BinSearch(arr []QQ, data int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2
		if data < arr[mid].QQuser {
			right = mid - 1
		} else if data > arr[mid].QQuser {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	path := "..\\QQ.txt"
	file, _ := os.Open(path)
	defer file.Close()
	br := bufio.NewReader(file)

	i := 0
	allStr := make([]QQ, N)
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		lineStr := string(line)
		lines := strings.Split(lineStr, "----")
		if len(lines) == 2 {
			allStr[i].QQuser, _ = strconv.Atoi(lines[0]) //Atoi(s string) (int, error)
			allStr[i].QQpass = lines[1]
		}
		i++
	}
	fmt.Println("印出在記憶體中")
	time.Sleep(time.Second)
	start := time.Now()
	fmt.Println("開始排序", len(allStr))
	allStr = QuickSortStruct(allStr)

	flag := true
	choose := ""
	for flag {
		fmt.Println("輸入查詢數據")
		var QQ int
		fmt.Scanf("%d\n", &QQ)
		// start := time.Now()
		index := BinSearch(allStr, QQ)
		if index == -1 {
			fmt.Println("找不到")
		} else {
			fmt.Println("找到", index, allStr[index].QQuser, allStr[index].QQpass)
		}
		fmt.Println("耗時", time.Since(start))
		fmt.Println("是否繼續查詢?(Y/N)")
		fmt.Scanln(&choose)
		if choose == "Y" {
			flag = true
		} else {
			break
		}
	}
}
