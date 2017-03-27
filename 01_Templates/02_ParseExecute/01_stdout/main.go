package main

import (
	"text/template"
	"log"
	"os"
)

func main()  {

	// need to add the relative path in case this is being run from Gogland ide
	// in case of running from command line remove the relative path
	relativePath := "01_Templates/02_ParseExecute/01_stdout/"

	tpl, err := template.ParseFiles(relativePath + "tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	er := tpl.Execute(os.Stdout, nil)
	if er != nil {
		log.Fatal("err in executing template")
	}
}
