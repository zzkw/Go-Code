package main

import "fmt"

func main() {
	s := "ab" +
		"cd"
	bs := []byte(s)
	bs[3] = 'D'
	fmt.Println(string(bs))
	u := "电脑"

	us := []rune(u) //中文rune
	us[1] = '话'
	fmt.Println(string(us))
}
