package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type people struct {
	Fname string
	Lname string
}

var tpl *template.Template
var arr []people
var firstname string
var lastname string

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/process", processHandler)
	http.HandleFunc("/action", actionHandler)
	http.ListenAndServe(":8086", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.FormValue("fname") != "" && r.FormValue("lname") != "" {
		firstname = r.FormValue("fname")
		lastname = r.FormValue("lname")
	} else {
		fmt.Println("Error Found")
	}
	a := people{Fname: firstname, Lname: lastname}
	arr = append(arr, a)
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "action.html", arr)
}
