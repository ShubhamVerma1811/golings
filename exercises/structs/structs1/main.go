// structs1
// Make me compile!

package main

import "fmt"

type Person struct {
	name string
	age  byte
}

func main() {

	person := Person{
		name: "Boris",
		age:  100,
	}

	fmt.Printf("Person %s and age %d", person.name, person.age)
}
