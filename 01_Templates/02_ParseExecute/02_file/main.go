package main

import (
	"text/template"
	"log"
	"os"
)

func main()  {

	// need to add the relative path in case this is being run from Gogland ide
	// in case of running from command line remove the relative path
	//relativePath := "01_Templates/02_ParseExecute/01_stdout/"
	relativePath := ""

	tpl, err := template.ParseFiles(relativePath + "tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	f, e := os.Create("index.html")
	defer f.Close()
	if e != nil {
		log.Fatal(e)
	}
	er := tpl.Execute(f, nil)
	if er != nil {
		log.Fatal("err in executing template")
	}
}
