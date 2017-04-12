package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
		Header      http.Header
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
	}

	fmt.Println(r.URL)
	tpl.ExecuteTemplate(w, "index.gohtml", data)

}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
