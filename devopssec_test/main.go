package main

import "fmt"

//constantes globale
const (
	tailleDamier  = 9
	symboleJoeur1 = "x"
	symboleJoeur2 = "0"
)

//varialble globale

var (
	tableauMorpion = [tailleDamier]string{
		"1", "2", "3",
		"4", "4", "6",
		"7", "8", "9"}

	joueur1 = true // Joueur 1 commence

)

//Afficher tableau
func affichage() {

	for i := 0; i < len(tableauMorpion); i++ {

		fmt.Print(" ", tableauMorpion[i], " ")
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}

}

//function principale

func main() {

	affichage()
}
