package jsonhandling

import (
	"encoding/json"
	"fmt"
	"os"
)

// struct for todo element
type Todo struct {
	Item      string `json:"item"`
	Category  string `json:"category"`
	Completed bool   `json:"completed"`
	Id        int    `json:"id"`
}

// global functions
func ListAll(fileName string) {
	todos := readJSON(fileName)
	printOut(todos)

}

func AddItem(title string, desc string, filename string) {
	//get list of exsisting items from json
	todos := readJSON(filename)

	//add the item
	if title == "" {
		fmt.Println("error: item title empty :" + title)
		return
	}

	id := todos[len(todos)-1].Id + 1
	item := Todo{title, desc, false, id}
	todos = append(todos, item)
	fmt.Println("Item has been added!")

	saveToJson(todos, filename)
}

func DeleteItem(id int, title string, filename string) {
	if id <= 1 {
		fmt.Println("error: please specify id example: \"delete -id 2\"")
		return
	}
	id = id - 1 //convert to zero start id
	fmt.Println("deleteing item")
	todos := readJSON(filename)
	todos = append(todos[:id], todos[id+1:]...)

	saveToJson(todos, filename)
}

//-----
//private functions

func saveToJson(list []Todo, filename string) {
	//save each todo item in json file

	// Open a new file for writing
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	file.Write([]byte("[")) //open bracket a start of file
	for i := 0; i < len(list); i++ {
		// Encode the data (write to the file)
		if i > 0 {
			file.Write([]byte(","))
		}

		list[i].Id = i + 1 //resets the id incase an item was deleted in the middle
		err = encoder.Encode(list[i])
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
		}
	}

	file.Write([]byte("]")) //close bracket at end of file
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
		fmt.Printf(" Id: %d - Item: %s - Category: %s - Completed: %t \n", s.Id, s.Item, s.Category, s.Completed)
	}
}
