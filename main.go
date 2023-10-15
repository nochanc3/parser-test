package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	City      string   `json:"city"`
	Email     string   `json:"email"`
	Interests []string `json:"interests"`
}

func main() {
	data := `
	{
	  "name": "",
	  "age": ,
	  "city": "",
	  "email": "@example.com",
	  "interests": [",,,"]
	}
	`

	var person Person

	err := json.Unmarshal([]byte(data), &person)
	if err != nil {
		fmt.Println("JSON error", err)
		return
	}

	fmt.Println("Results:")
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("City: %s\n", person.City)
	fmt.Printf("Email: %s\n", person.Email)
	fmt.Printf("Hobby: %v\n", person.Interests)
}
