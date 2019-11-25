package main

import "fmt"

func main() {
	if a, b := 1, 2; a >= 1 { //GOif语句节选
		fmt.Println(a)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Printf("hello, world!你好，世界！")
}
