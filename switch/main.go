package main

import "fmt"

func main() {
	weekday := 6 // 1 == monday, 7 == sunday
	fmt.Printf("Day of the weekday=%d. What's for today ?\n", weekday)

	switch weekday {
	case 1:
		fmt.Println("Beginning of the week, let's get to work")

	case 3:
		fmt.Println("wednestday, the half is done")

	//possibilité de definir plusieurs choix
	case 6, 7:
		fmt.Println("It's the week-end")

	default:
		fmt.Println("Nothing special about this day")
	}

	hour := 10
	fmt.Printf("Current time=%d\n ", hour)
	switch {
	case hour < 12:
		fmt.Println("It' the morning")

	case hour > 12 && hour < 18:
		fmt.Println("It's the afternoom")
	default:
		fmt.Println("It's the evening")
	}
}
