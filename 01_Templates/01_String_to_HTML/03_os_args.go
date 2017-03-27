package main

import (
	"fmt"
	"os"
	"io"
	"strings"
)

func main()  {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	name := os.Args[1]

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
