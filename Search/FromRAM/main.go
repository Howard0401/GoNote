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
	i := 0
	br := bufio.NewReader(file)

	//Assign value to slice

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

	fmt.Println("請輸入要查詢的數據")
	var inputStr string = "yinchung"
	fmt.Scanln(inputStr)

	startTime := time.Now()
	for j := 0; j < N; j++ {
		if strings.Contains(fileStrings[j], inputStr) {
			fmt.Println(fileStrings[j])
		}
	}
	fmt.Println("共用了", time.Since(startTime))

}
