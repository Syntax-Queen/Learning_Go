package main

import "fmt"

// import "fmt"

func main() {
	
	var greeting string //zero-value is an empty string ""

	greeting = "Hello world!"

	fmt.Println(greeting)

	var count int
	count = 10

	fmt.Println(count)

	var isRunning bool
	isRunning = true

	fmt.Println(isRunning)


	var firstName, lastName string
	firstName = "Queen"
	lastName = "Adebisi"

	fmt.Println(firstName, lastName)

	email := "test@gmail.com" //when you want to declare and initialize your variable at the same time

	fmt.Println(email)

	age := 24
	fmt.Println(age)

}