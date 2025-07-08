package main

import "fmt"

func main() {
	fmt.Println("Fourth Lesson A.11 Constant")

	const firstName string = "Jhon"

	fmt.Print("Hello ", firstName, "!\n")

	fmt.Println("jhon wick")
	fmt.Println("jhon", "wick")

	fmt.Print("jhon wick\n")
	fmt.Print("jhon ", "wick\n")
	fmt.Print("jhon", " ", "wick\n")

	const (
		square         = "Kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	const (
		a = "Constant"
		b
		c
	)

	const one, two = 1, 2

	fmt.Printf("this is a const %s %t %d %.2f \n", square, isToday, numeric, floatNum)
	fmt.Printf("this is const %s %s %s \n", a, b, c)
	fmt.Printf("this is a const %d %d \n", one, two)
}
