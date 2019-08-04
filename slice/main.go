package main

import "fmt"

// Un slice représente une tranche d'un tableau
// un slice est dynamique
// syntaxe :
// s := make([]type, taille, capacité)
//taill: nombre d'élement du slice
//capacite: nombre d'élement du tableau

// %v pour les tableaux

func main() {
	numerique := make([]int, 2, 3)
	numerique[0] = 4
	numerique[1] = 5
	fmt.Printf("%v longueur=%d, capacite=%d)\n", numerique, len(numerique), cap(numerique))

	//ajout d'un chiffre au tableau "numerique". ici nous sommes toujours dans la capcité de 3
	numerique = append(numerique, 64)
	fmt.Printf("%v longueur=%d, capacite=%d)\n", numerique, len(numerique), cap(numerique))

	// ici on dépasse la capacité initiale, donc go va augmenté la capacité de facon auto.
	numerique = append(numerique, 120)
	fmt.Printf("%v longueur=%d, capacite=%d)\n", numerique, len(numerique), cap(numerique))

	fmt.Println("****************************************")
	// syntaxe courte de slice
	lettres := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("%v longueur=%d, capacite=%d)\n", lettres, len(lettres), cap(lettres))

	//ajout d'un lettre
	lettres = append(lettres, "!")
	fmt.Printf("%v longueur=%d, capacite=%d)\n", lettres, len(lettres), cap(lettres))

}
