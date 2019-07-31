package main

import (
	"fmt"
)

func printInfoNoParam() {
	fmt.Printf("Name=%s, age=%d, email=%s\n", "bob", 10, "bob@yahoo.fr")

}

func printInfoParam(name string, age int, email string) {
	fmt.Printf("Name=%s, age=%d, email=%s\n", name, age, email)

}

func avg(x, y float64) float64 {
	return (x + y) / 2
}

func sumNameReturn(x, y, z int) (sum int) { //ici on déclare la variable dans la fonction

	sum = x + y + z
	return sum
}

func main() {

	printInfoNoParam()

	printInfoParam("Philippe", 42, "pdevezeaud@yahoo.fr")

	resultat := avg(16.0, 16.0)
	fmt.Printf("Average resultat =%f\n", resultat)

	somme := sumNameReturn(2, 2, 2)
	fmt.Printf("la somme est = %d\n", somme)

}
