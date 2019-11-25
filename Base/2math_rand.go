package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("a number form 1-100", rand.Intn(100)) //随机数，在此GO程序中出现一个随机数后重新运行还是那个数
}
func main() {
	foo()
	fmt.Println("4 squrt root is", math.Sqrt(4)) //开平方
}
