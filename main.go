package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/stuff", HandleStuff)
	fmt.Println("Listening on :8041")
	http.ListenAndServe(":8041", nil)
}
