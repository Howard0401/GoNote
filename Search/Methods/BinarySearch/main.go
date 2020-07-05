package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type input struct {
	user     string
	md5      string
	email    string
	password string
}

func QuickSort(arr []input) []input {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		spliteData := arr[0]
		low := make([]input, 0)
		high := make([]input, 0)
		mid := make([]input, 0)
		mid = append(mid, spliteData)

		for i := 1; i < length; i++ {
			if arr[i].user < spliteData.user {
				low = append(low, arr[i])
			} else if arr[i].user > spliteData.user {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

func BinSearch(arr []input, data string) int {
	low := 0
	high := len(arr) - 1
	for high >= low {
		mid := (high + low) / 2
		if arr[mid].user > data {
			high = mid - 1
		} else if arr[mid].user < data {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	//Read file to buffer
	path := "..\\..\\uuu9.com.sql"
	file, _ := os.Open(path)
	br := bufio.NewReader(file)

	//Assign DataN
	allData := make([]input, 0)
	i := 0
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		lineStr := string(line)
		lines := strings.Split(lineStr, " | ")
		if len(lines) == 4 {
			in := input{
				user:     lines[0],
				md5:      lines[1],
				email:    lines[2],
				password: lines[3],
			}
			allData = append(allData, in)
		}
		i++
	}
	fmt.Println("載入完成")
	allData = QuickSort(allData)
	fmt.Println("排序完成")
	flag := true
	for flag {
		fmt.Println("請輸入欲查詢的用戶名") //ex.haiyu3399
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入
		starttime := time.Now()
		// for i := 0; i < len(allData); i++ {
		// 	if allData[i].user == inputstr {
		// 		fmt.Println(allData[i])
		// 	}
		// }
		index := BinSearch(allData, inputstr) //用Binary取代linear
		fmt.Println("本次查询用了", time.Since(starttime))
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
