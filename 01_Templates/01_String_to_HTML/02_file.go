package main

import (
	"os"
	"io"
	"fmt"
	"strings"
)

func main()  {
	name := "Rohan"

	str := fmt.Sprint(`<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World !</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>`)


	f, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, strings.NewReader(str))
}
