package main

import "fmt"

func main() {
	fmt.Println("Fifth Lesson A.12 Operator")

	var value uint16 = ((12*12)/12 + (1 * 1))

	var isEqual bool = (value == 13)

	fmt.Printf("The value is %d \n", value)
	fmt.Printf("The value is Equal to %d = %t \n", value, isEqual)

	var left bool = false
	var right bool = true

	var leftAndRight bool = left && right
	var leftOrRight bool = left || right
	var leftReverse bool = !left

	fmt.Printf("Left and Right %t , Left or Right %t, Left Reverse %t \n", leftAndRight, leftOrRight, leftReverse)

}
