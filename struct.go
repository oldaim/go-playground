package main

import "fmt"

func main() {

	newPerson := CommonPerson{
		Name: "kim",
		Age:  0,
	}

	fmt.Printf("name is %s\n", newPerson.getName())
	fmt.Printf("Age is %d", newPerson.getAge())

}

type Person interface {
	getName() string
	getAge() int
}
type CommonPerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (receiver CommonPerson) getName() string {
	return receiver.Name
}

func (receiver CommonPerson) getAge() int {
	return receiver.Age
}
