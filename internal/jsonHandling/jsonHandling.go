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

func ListAll(fileName string) {
	todos := readJSON(fileName)
	printOut(todos)

}

//-----

func saveToJson(list []Todo, filename string) {
	//save each todo item in json file
	file, _ := os.OpenFile(filename, os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(list[0])
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

func AddItem(title string, desc string, filename string) {
	//get list of exsisting items from json
	fmt.Println("adding..")
	todos := readJSON(filename)

	//add the item
	if title == "" {
		fmt.Println("error: item title empty :" + title)
		return
	}

	id := todos[len(todos)-1].Id + 1
	item := Todo{title, desc, false, id}
	todos = append(todos, item)
	fmt.Println(todos)
	//save json file (overwrite with new list)

	saveToJson(todos, filename)
}

func printOut(todoSlice []Todo) {
	for _, s := range todoSlice {
		fmt.Printf(" Id: %d - Title: %s - Desc: %s - Completed: %t \n", s.Id, s.Title, s.Desc, s.Completed)
	}
}
