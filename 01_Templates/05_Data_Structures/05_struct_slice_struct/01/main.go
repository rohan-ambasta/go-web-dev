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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Player    []player
	Transport []car
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

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	players := []player{kohli, rahane, ashwin, jadeja, umesh}

	cars := []car{f, c}

	data := items{
		Player:    players,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
