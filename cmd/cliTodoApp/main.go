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
	var id *int

	if os.Args[1] == "add" {
		//if flag has add option make a new flag set that looks at the args after flag add
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title = addCmd.String("title", "", "a string")
		desc = addCmd.String("c", "personnal", "category")
		addCmd.Parse(os.Args[2:])
	}
	if os.Args[1] == "delete" {
		//if flag has delete option make a new flag set that looks at the args after flag delete
		removeCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id = removeCmd.Int("id", -1, "item id to delete")
		removeCmd.Parse(os.Args[2:])
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
	case "delete":
		jsonhandling.DeleteItem(*id, "", fileName)
	}

}
