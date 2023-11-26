package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi bardia gopher!")
}

func main() {
	http.HandleFunc("/test", index)
	http.ListenAndServe(":4232", nil)
}
