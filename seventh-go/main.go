package main

import "fmt"

func main() {
	fmt.Println("Lesson Seventh A.14 Looping")

	// for i := 1; i <= 5; i++ {
	// 	fmt.Printf("Number : %d \n", i)
	// }

	// var j uint8 = 0

	// for j < 5 {
	// 	fmt.Printf("Number : %d \n", j)
	// 	j++
	// }

	// var k uint = 0

	// for {
	// 	fmt.Printf("Number : %d \n", k)
	// 	k++
	// 	if k == 5 {
	// 		break
	// 	}
	// }

	// var xs string = "123" // string
	// for i, v := range xs {
	// 	fmt.Println("index=", i, "Value=", v)
	// }

	// var ys = [5]int{10, 20, 30, 40, 50} // array
	// for _, v := range ys {
	// 	fmt.Println("Value=", v)
	// }

	// var zs = ys[0:2] // slice
	// for _, v := range zs {
	// 	fmt.Println("Value=", v)
	// }

	// var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} //map
	// for k, v := range kvs {
	// 	fmt.Println("Key=", k, "Value=", v)
	// }

	// for range kvs {
	// 	fmt.Println("Done")
	// }

	// for i := range 5 {
	// 	fmt.Print(i)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Number : ", i)
	// }

outer:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 4 {
				break outer
			}
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

}
