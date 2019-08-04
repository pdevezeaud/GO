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
	// ici on va faire des tranche de slice avec des slice
	// on stocke de sub1 un morceau de lettres dans sub1 pui dans sub2
	// ici on prend de l'indice 0 à deux soit gol pur un lang pour l'autre
	fmt.Println("********************************************")
	// on peut ommettre aussi le 0
	sub1 := lettres[0:2]
	fmt.Printf("%v longueur=%d, capacite=%d)\n", sub1, len(sub1), cap(sub1))

	fmt.Println("********************************************")
	// ici go autorise de ne rien mettre,implicitement il en découle qu'il va jusqu'à la fin du tableau
	sub2 := lettres[2:]
	fmt.Printf("%v longueur=%d, capacite=%d)\n", sub2, len(sub2), cap(sub2))

	fmt.Println("********************************************")
	fmt.Println("********************************************")
	sub2[3] = "up"
	fmt.Printf("%v longueur=%d, capacite=%d)\n", lettres, len(lettres), cap(lettres))

	fmt.Println("**********  COPIE DE SLICE  **********************")
	fmt.Println("************                **********************")
	// on initialise la variable dans laquel on va stocké la copie
	subCopy := make([]string, len(sub1))
	copy(subCopy, sub1)
	fmt.Printf("copy %v longueur=%d, capacite=%d)\n", subCopy, len(subCopy), cap(subCopy))

}
