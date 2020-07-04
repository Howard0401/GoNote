package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const N = 84331469

//使用Struct輔助搜尋

type QQ struct {
	QQuser int
	QQpass string
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

	flag := true
	choose := ""
	for flag {
		fmt.Println("輸入查詢數據")
		var QQ int
		fmt.Scanf("%d\n", &QQ)
		start := time.Now()
		for j := 0; j < N; j++ {
			if allStr[j].QQuser == QQ {
				fmt.Println(j, allStr[j].QQuser, allStr[j].QQpass)
			}
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
