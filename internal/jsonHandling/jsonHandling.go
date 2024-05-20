package jsonhandling

import (
	"encoding/json"
	"fmt"
	"os"
)

// struct for todo element
type Todo struct {
	Title     string
	Desc      string
	Completed bool
	Id        int
}

func SaveToJson(list []Todo) {
	//save each todo item in json file

}

func readJSON(fileName string) []Todo {

	var todo []Todo

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
	}

	err = json.Unmarshal(file, &todo)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return []Todo{}
	}

	return todo
}

func printOut(todoSlice []Todo) {
	for _, s := range todoSlice {
		fmt.Printf(" Id: %d - Title: %s - Desc: %s - Completed: %t \n", s.Id, s.Title, s.Desc, s.Completed)
	}
}

func ListAll(fileName string) {
	todos := readJSON(fileName)
	printOut(todos)

}
