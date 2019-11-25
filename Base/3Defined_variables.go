package main

import (
	"fmt"
)

var x, y, z int       //函数外定义变量
var s, n = "abc", 123 //函数外定义变量并赋值(可忽略变量类型，Go直接根据赋值参数给予变量类型)
var (                 //定义不同类型变量
	a int
	b float32
)

func main() {
	n, s := 0x1234, "Helloman" //函数内定义变量并赋值(最简版)
	fmt.Println(x, s, n)
}
