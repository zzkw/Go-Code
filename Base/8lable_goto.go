package main

import (
	"fmt"
)

func main() {

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL1
			}
			fmt.Println(i)
		}
	}
LABEL1:
	fmt.Println("da")
}
