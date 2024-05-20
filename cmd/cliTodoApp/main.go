package main

import (
	"flag"
	"fmt"

	jsonhandling "github.com/samL23/cliTodoApp/internal/jsonHandling"
)

func main() {
	// fetch the previous todo elements from json files
	// add todo element into a slice
	fileName := "data.json"
	flag.Parse()
	fmt.Println(flag.Arg(0))

	switch flag.Arg(0) {
	case "":
		//error print help
	case "list":
		jsonhandling.ListAll(fileName)
	}

}
