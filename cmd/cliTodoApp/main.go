package main

import (
	"encoding/json"
	"fmt"
	"os"

	jsonhandling "github.com/samL23/cliTodoApp/internal/jsonHandling"
)

// struct for todo element
type Todo struct {
	Title string
	Desc  string
	Id    int
}

func main() {
	// fetch the previous todo elements from json files
	// add todo element into a slice
	testTodo1 := Todo{Title: "go shopping", Desc: "buy papers", Id: 1}
	var todos []Todo

	todos = readJSON("data.json")
	fmt.Print(todos)
	todos = append(todos, testTodo1)
	fmt.Println(todos)
	fmt.Println(jsonhandling.TestingInternalPkg())

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
