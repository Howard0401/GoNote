package main

import "fmt"

func main() {
	slice := []int{0, 3, 6, 5, 7}
	testMap1 := make(map[int]*int)
	testMap2 := make(map[int]int)
	count := 0
	for k, v := range slice { //part1
		//part2
		fmt.Printf(" key: %v, key address: %v, value: %v, value_address: %v\n", k, &k, v, &v)
		//part3
		testMap1[k] = &v
		testMap2[k] = v
		count += 1
		//part4
		slice = append(slice, v)
	}
	fmt.Println("now slice is:", slice, "count=", count)

	fmt.Println("testMap1:")
	for k, v := range testMap1 {
		fmt.Printf("%d => %d\n", k, *v)
	}

	fmt.Println("testMap2:")
	for k, v := range testMap2 {
		fmt.Printf("%d => %d\n", k, v)
	}
}
