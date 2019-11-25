package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/xml"
)

func main() {
	resp, _ := http.Get("https://zzkw.github.io/")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}
