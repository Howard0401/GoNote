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

func main() {
	path := "..\\..\\uuu9.com.sql"
	file, _ := os.Open(path)
	br := bufio.NewReader(file)
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
	fmt.Println(i)

	flag := true
	for flag {
		fmt.Println("請輸入欲查詢的用戶名")
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入
		starttime := time.Now()
		for i := 0; i < len(allData); i++ {
			if allData[i].user == inputstr {
				fmt.Println(allData[i])
			}
		}
		fmt.Println("本次查询用了", time.Since(starttime))
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
