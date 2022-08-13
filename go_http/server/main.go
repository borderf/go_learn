package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/boy", BoyHandler)
	http.HandleFunc("/girl", testGirl)
	http.ListenAndServe(":8088", nil)
}

func BoyHandler(w http.ResponseWriter, r *http.Request) {
	io.Copy(os.Stdout, r.Body)
	fmt.Println("request headers: ")
	for k, v := range r.Header {
		fmt.Printf("%s = %v\n", k, v)
	}
	fmt.Println("request cookies: ")
	for _, cookie := range r.Cookies() {
		fmt.Printf("%s = %v\n", cookie.Name, cookie.Value)
	}
	fmt.Fprint(w, "hello boys")
}

func testGirl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello girls")
}
