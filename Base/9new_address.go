package main

import (
	"fmt"
)

func main() {
	a := [...]int{9: 4}
	fmt.Println(a)
	var p *[10]int = &a //指向数组的指针(是指针)
	fmt.Println(p)
	n := new([10]int)
	fmt.Println(n)
}
