package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", basicCheck)
	http.HandleFunc("/check/", check)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)

}

func check(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from Rest APi")

}

func basicCheck(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from Backend")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "monisha")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
