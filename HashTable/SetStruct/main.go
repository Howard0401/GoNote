package main

import "fmt"

type Set struct {
	buf  []interface{}
	num  int
	hash map[interface{}]bool
}

//建立一個可以變長的Set
func NewSet() *Set {
	return &Set{make([]interface{}, 0), 0, make(map[interface{}]bool)}
}

func (this *Set) Add(value interface{}) bool {
	if this.isExist(value) {
		return false
	} else {
		this.buf = append(this.buf, value)
		this.num++
		this.hash[value] = true
		return true
	}
}

func (this *Set) isExist(value interface{}) bool {
	return this.hash[value]
}

func (this *Set) Strings() []interface{} {
	return this.buf
}

func main() {
	set := NewSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	fmt.Println(set)
}
