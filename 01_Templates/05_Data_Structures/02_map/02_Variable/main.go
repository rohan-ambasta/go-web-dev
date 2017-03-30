package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

// *****************************************
// You get to pass in one value - that's it!
// *****************************************
func main() {
	values := map[string]string{
		"captain":          "kohli",
		"vice-captain":     "rahane",
		"off-spinner":      "ashwin",
		"left-arm-spinner": "jadeja",
		"fast-bowler":      "umesh"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", values)
	if err != nil {
		log.Fatalln(err)
	}

}
