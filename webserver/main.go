package main

import (
	"fmt"
	"net/http"
	
)

func main() {

	http.HandleFunc("/hello", helloHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)
	http.ListenAndServe(":8080", nil)
}
func helloHandleFunc(k http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(k,"r.URL.Path:,%s",r.URL.Path)
	fmt.Fprint(k, "yo yo yo yo yo yo yo happy")
}
func aboutHandleFunc (k http.ResponseWriter, r *http.Request){
	fmt.Fprint(k, "happy to you brother where are you")

	
}
