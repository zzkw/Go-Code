package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(&s) //s的地址

	s, y := "hello", 20 // 重新赋值（地址还是它）: 与前 s 在同一层次的代码块中，且有新的变量被定义。
	fmt.Println(&s, y)  // 通常函数多返回值 err 会被重复使用。

	{
		s, z := 1000, 30 // 定义新同名变量（新地址存放新变量）: 不在同一层次代码块。
		fmt.Println(&s, z)
	}
}
