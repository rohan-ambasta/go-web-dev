package main

import (
	"log"
	"os"
	"strings"
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

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 3 {
		s = s[:3]
	}
	return s
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

	data := struct {
		Player    []player
		Transport []car
	}{
		Player:    players,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
