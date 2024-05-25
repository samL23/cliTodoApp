package main

import (
	"flag"
	"os"

	jsonhandling "github.com/samL23/cliTodoApp/internal/jsonHandling"
)

func main() {
	const fileName = "output.json"
	// fetch the previous todo elements from json files
	// add todo element into a slice
	var title *string
	var desc *string

	if os.Args[1] == "add" {
		//if flag has add option make a new flag set that looks at the args after flag add
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title = addCmd.String("title", "root", "a string")
		desc = addCmd.String("c", "personnal", "category")
		addCmd.Parse(os.Args[2:])
	}

	//parse original full flag
	flag.Parse()

	switch flag.Arg(0) {
	case "":
		//error print help
	case "list":
		jsonhandling.ListAll(fileName)
	case "add":
		jsonhandling.AddItem(*title, *desc, fileName)
	}

}
