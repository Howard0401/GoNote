package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 10
	i := <-ch
	fmt.Println(i)
}