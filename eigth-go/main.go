package main

import "fmt"

func main() {
	fmt.Println("Lesson Eight A.15 Array")

	var names [4]string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	for i := range fruits {
		fmt.Printf("element %d \n", i)
	}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Elemen %d : %s \n", i, fruits[i])
	// }

	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("element \t", fruits)

	var numbers = [...]int{1, 2, 3, 4, 5}

	fmt.Println("data array : \t", numbers)

	var numbers1 = [2][3]int{[3]int{3, 2, 2}, [3]int{1, 2, 3}}
	var numbers2 = [2][3]int{{3, 2, 3}, {1, 2, 3}}

	fmt.Println(numbers1)
	fmt.Println(numbers2)

	var fruitsTwo = make([]string, 2)

	fruitsTwo[0] = "apple"
	fruitsTwo[1] = "Banana"
}
