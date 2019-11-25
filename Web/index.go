package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>whoa,Go is neat!</h1>
<p>Go let's do it!</p>`)
}
func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
