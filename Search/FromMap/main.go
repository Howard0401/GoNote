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

func main() {
	allStr := make(map[int]string, N)
	path := "..\\QQ.txt"
	file, _ := os.Open(path)
	defer file.Close()

	br:=bufio.NewReader(file)
	
	i:=0
	for {
		line,_,end :=br.ReadLine()
		if end == io.EOF{
			break
		}
		lineStr:=string(line)
		lines:=strings.Split(lineStr,"----")
		if len(lines)==2 {
			User,_:=strconv.Atoi(lines[0])
			Pass:=lines[1]
			allStr[User]=Pass
		}
		i++
	}
	fmt.Println("印出在記憶體中")
	time.Sleep(time.Second)
	flag := true
	for flag {
		fmt.Println("輸入想查的值")
		var QQ int 
		fmt.Scanf("%d\n",&QQ)
		QQPass,err:= allStr[QQ]
		if err {
			fmt.Println(QQ,QQPass,"存在")
		}else{
			fmt.Println(QQ,"不存在")
		}

		fmt.Println("是否繼續查詢?")
		var choose string
		fmt.Scanf("%s\n",&choose)
		if choose == "Y"{
			flag = true
		}else {
			break
		}
	}
}