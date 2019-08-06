//  range est une fonction et prend en source de donnée des data ( renvoi le int qui sera la valeur de l'index )

package main

import "fmt"

func main() {
	noms := []string{"toto", "titi", "tata"}
	// ici i sera la variable qui va iterer sur l'index

	for i, noms := range noms {
		fmt.Printf("username=%s (index=%d)\n", noms, i)
	}

	// non utilisation de l'index
	for _, caractere := range "DEVEZEAUD" {
		fmt.Printf("%v\n", string(caractere))
	}
}
