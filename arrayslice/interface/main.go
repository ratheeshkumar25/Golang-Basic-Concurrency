package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof"
}

func (c Cat) Speak() string {
	return "Meow"
}

func main() {
	var myPet Animal
	myPet = Dog{Name: "Buddy"}
	fmt.Println(myPet.Speak())

	myPet = Cat{Name: "Cabbi"}
	fmt.Println(myPet.Speak())
}
