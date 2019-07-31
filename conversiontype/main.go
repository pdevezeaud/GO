package main

import (
	"fmt"
)

func main() {

	//ici on mais un f car c'est un float
	var percentage float64 = 2.0

	//+= est identique à percentage = percentage + 31
	percentage += 34

	fmt.Printf("Current percentage %f%%\n", percentage)

	//ici on fait la conversion en precisant que l'on veut une nouvelle donnée en int
	fmt.Printf("Int valeur=%d%%\n", int(percentage))

	n := 42
	// ici on converti 42 pour qu'il soit intégré dans un float
	f := float64(n) + 0.42
	fmt.Printf("float value= %f\n", f)

}
