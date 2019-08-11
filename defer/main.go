package main

import "fmt"

// defer (repoussé l'execution à plus tard)

func start() {
	fmt.Println("Start")
}

func finish() {
	fmt.Println("Finish")
}

func main() {

	start()
	defer finish() //LIFO = last in First Out
	//finish()
	// Goodbye bob
	// Goodbye bobette
	// Goodbye John
	// il commence par john, puis bobette, puis bob et lance le finish

	noms := []string{"Bob", "Bobette", "John"}

	for _, noms := range noms {
		fmt.Printf("Hello %v\n", noms)
		defer fmt.Printf("Goodbye %v\n", noms)
	}

}
