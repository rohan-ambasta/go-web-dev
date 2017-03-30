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
	rahane := player{
		Name: "Ajinkya Rahane",
		Role: "Vice-Captain",
	}
	ashwin := player{
		Name: "Ravi Ashwin",
		Role: "Off spinner",
	}
	jadeja := player{
		Name: "Ravindra Jadeja",
		Role: "Left arm spinner",
	}
	umesh := player{
		Name: "Umesh Yadav",
		Role: "Fast Bowler",
	}

	players := []player{kohli, rahane, ashwin, jadeja, umesh}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", players)
	if err != nil {
		log.Fatalln(err)
	}

}
