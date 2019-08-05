package main

import "fmt"

func main() {

	sum := 0
	for i := 1; i <= 10; i++ {
		println("boucle loop =", i)
		sum = sum + i

	}
	fmt.Printf("Sum result =%d\n", sum)

	fmt.Println("*************************************************")
	eventCount := 0
	for eventCount < 3 {
		fmt.Println("Recuperation d'evenement")
		eventCount++
		if eventCount == 3 {
			fmt.Printf("Got %d notification, upgrade du telephone\n", eventCount)
		}

		i := 0
		for {
			i++
			if i%2 != 0 {
				fmt.Println("Odd Looping")
				continue
			}
			fmt.Println("Looping... Paire")

			if i >= 10 {
				fmt.Println("Loop End")
				break
			}

		}

	}
}
