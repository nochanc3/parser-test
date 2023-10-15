package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	City      string   `json:"city"`
	Email     string   `json:"email"`
	Interests []string `json:"interests"`
}

func main() {
	// Read data from the "data.json" file
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	var person Person

	// Parse the JSON data
	err = json.Unmarshal(data, &person)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Parsing result:")
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("City: %s\n", person.City)
	fmt.Printf("Email: %s\n", person.Email)
	fmt.Printf("Interests: %v\n", person.Interests)
}
