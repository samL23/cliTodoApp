package jsonhandling

import (
	"encoding/json"
	"fmt"
	"os"
)

// struct for todo element
type Todo struct {
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
	Id        int    `json:"id"`
}

func ListAll(fileName string) {
	todos := readJSON(fileName)
	printOut(todos)

}

//-----

func saveToJson(list []Todo, filename string) {
	//save each todo item in json file

	// Open a new file for writing
	file, err := os.Create("output.json")
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
