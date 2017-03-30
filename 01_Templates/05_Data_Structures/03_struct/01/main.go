package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type player struct {
	Name string
	Role string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

// *****************************************
// You get to pass in one value - that's it!
// *****************************************
func main() {

	kohli := player{
		Name: "Virat Kohli",
		Role: "Captain",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", kohli)
	if err != nil {
		log.Fatalln(err)
	}

}
