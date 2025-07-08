package main

import "fmt"

func main() {

	fmt.Println("Second Lesson A.9 Variable")

	_ = "This is predifined varialbe, is ok to not use it"

	var firstName string = "jhon"
	lastName := "Wick"

	var first, second, third string

	first, second, third = "one", "two", "three"

	fmt.Printf("%s %s %s \n", first, second, third)

	fmt.Printf("Hello %s %s!\n", firstName, lastName)
}
