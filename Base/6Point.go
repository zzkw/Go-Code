package main

import (
	"fmt"
)

func main() {
	var a int = 1
	var b *int = &a  //*b=*&a=a, b=&a(a的地址)
	var c **int = &b //定义**c=**&b=*b=*&a=a,*c=b=&a,c=&b(b的地址)，&c(c的地址)
	var x int = *b   //x=*b=a
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*b = ", *b)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("**c = ", **c)
	fmt.Println("c = ", c)
	fmt.Println("&c = ", &c)
	fmt.Println("x = ", x)
	x, y := 1, 2
	ar := [...]*int{&x, &y} //指针数组(是数组)
	fmt.Println(ar)
}
