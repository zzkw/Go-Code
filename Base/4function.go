package main

import "fmt"

func test() {
	data, _, i := [3]int{0, 1, 2}, 2, 0 //函数内部:=定义变量数组data[0]，data[1],data[2]并赋值0,1，2，且i=0
	i, data[i] = 2, 100                 // (data[0] = 100),(i = 0) -> (i = 2) | 从右往左赋值
	fmt.Println(i, data[i], data[0])
}
func main() {
	test()
}
