package main

import (
	"html/template"
	"log"
	"net/http"
)

func formPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	if r.Method == "POST" {
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		age := r.FormValue("age")
		data = map[string]interface{}{"first": "Hello, my full name is", "second": "I am ", "fname": fname, "lname": lname, "age": age}
		tmpl, _ := template.ParseFiles("static/formFinal.html")
		tmpl.Execute(w, data)
		return

	}
	tmpl, _ := template.ParseFiles("static/form.html")
	tmpl.Execute(w, nil)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formPage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
