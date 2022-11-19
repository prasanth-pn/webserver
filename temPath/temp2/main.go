package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

var tp *template.Template

func main() {
	//tp1, _ = template.ParseFiles("index1.html")
	//tp1, _ = template.ParseFiles("data1/index2.html")
	//tp1, _ = template.ParseFiles("data1/data2/index3.html")
	//tp1, _ = template.ParseFiles("../index4.html")
	//func (t Template)ParseFiles(filenames ...string) (*Template,error)
	tp,_=tp.ParseFiles("index1.html")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tp.Execute(w, nil)
	//fmt.Println(r)
}
