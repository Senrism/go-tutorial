package main

import "fmt"

func main() {
	fmt.Println("Third Lesson A.10 Data type")

	var positiveNumber uint8 = 89
	var negativeNumber = -123456789
	var decimalNumber = 2.69
	var exist bool = true
	var message string = `this is a message with backtick so i can use \ / without getting escaped`

	fmt.Printf("positive value %d \n", positiveNumber)
	fmt.Printf("negative value %d \n", negativeNumber)
	fmt.Printf("decimal value %f \n", decimalNumber)
	fmt.Printf("decimal value less 0 after comma %.2f \n", decimalNumber)
	fmt.Printf("exist? %t \n", exist)
	fmt.Printf("the message is %s \n", message)
}
