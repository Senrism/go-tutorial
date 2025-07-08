package main

import "fmt"

func main() {
	fmt.Println("Sixth Lesson A.13 Conditioning")

	var point uint = 8

	if point == 10 {
		fmt.Println("Perfect Point")
	} else if point >= 7 && point <= 8 {
		fmt.Println("Good point")
	} else if point <= 7 && point >= 5 {
		fmt.Println("Ya so so la")
	} else {
		fmt.Println("Badddddd")
	}

	var nextPoint = 8840.0

	if percent := nextPoint / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect !\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good !\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad \n", percent, "%")
	}

	var switchPoint = 8

	switch switchPoint {
	case 8:
		fmt.Println("Hey that's pretty good")
		fallthrough
	case 7, 6, 5, 4:
		fmt.Println("That's ok")
	default:
		{
			fmt.Println("Aigoo")
		}
	}
}
