package main

import (
	"html/template"
	"net/http"
)

var tp *template.Template
var name="prasanth"

func main() {
	//func ParseGlob(pattern string)(*Template,error)
	//tp, _ = template.ParseGlob("templates/*html")
	// func (t *Template)ParseFiles(filenames ...string)(*Template,error)
	tp, _ = tp.ParseGlob("templates/*.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about",aboutHandler)
	http.HandleFunc("/contact",contactHandler)
	http.HandleFunc("/home",homeHandler)
	http.HandleFunc("/login",loginHandler)
	http.ListenAndServe(":8080",nil)


}
func indexHandler(w http.ResponseWriter,r *http.Request){
	//func(t *Template)ExecuteTemplate(wr io.Writer,name string,data interface{})error
tp.ExecuteTemplate(w,"index.html",name)
}

func contactHandler(w http.ResponseWriter,r *http.Request){
	tp.ExecuteTemplate(w,"contact.html",nil)
}
func aboutHandler(w http.ResponseWriter, r *http.Request){
	tp.ExecuteTemplate(w, "about.html",nil)
}
func homeHandler( w http.ResponseWriter, r *http.Request){
	
		tp.ExecuteTemplate(w,"home.html",nil)
	
}
func loginHandler(w http.ResponseWriter, r *http.Request){
	tp.ExecuteTemplate(w,"login.html",nil)
}
