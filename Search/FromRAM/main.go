package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const N = 84331469

func main() {

	//Read the file
	fileStrings := make([]string, N)
	path := "..\\QQ.txt"
	file, _ := os.Open(path)
	defer file.Close()

	br := bufio.NewReader(file)

	//Assign value to slice
	i := 0
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		fileStrings[i] = string(line)
		i++
	}
	fmt.Println()
	fmt.Println("數據放入記憶體中", i, "行")
	time.Sleep(time.Second)

	//在試的時候有遇到一個問題，就是Debug Mode不能輸入
	// the VS Code debugger console doesn't currently support piping any input through to stdin.
	// https://github.com/Microsoft/vscode-go/issues/219
	flag := true
	var choose string
	var inputStr string = "yinchung"

	for flag {
		fmt.Println("請輸入要查詢的數據")
		fmt.Scanln(&inputStr)

		startTime := time.Now()
		for j := 0; j < N; j++ {
			if strings.Contains(fileStrings[j], inputStr) {
				fmt.Println(fileStrings[j])
			}
		}
		fmt.Println("共用了", time.Since(startTime))

		fmt.Println("是否繼續搜尋?(Y/N")
		fmt.Scanln(&choose)
		if choose == "Y" {
			flag = true
		} else {
			flag = false
		}
	}
}
