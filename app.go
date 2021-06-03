package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	name, path, err := CreateStudent()

	if err == nil {
		fmt.Println("..........Creating Document..........")
		fmt.Println("Document for ", name, "is added at", path)
	}

	getStudent(path)

}

func getStudent(path string) {

	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	data := Student{}

	if err := json.Unmarshal([]byte(file), &data); err != nil {
		fmt.Println(err)
	}

	fmt.Println("..........Reading Document..........")
	fmt.Println(data)

}
