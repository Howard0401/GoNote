package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	path := "..\\QQ.txt"
	file, _ := os.Open(path)
	count := 0
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		// fmt.Println(string(line)) //Check if read
		lineStr := string(line)
		if strings.Contains(lineStr, "yincheng") {
			// fmt.Println(lineStr)
			count++
		}
	}
	fmt.Println("共用了", time.Since(startTime), "一共", count, "行")
}
