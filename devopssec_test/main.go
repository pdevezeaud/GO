package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//constantes globales qui ne bourgerons pendant le deroulement du programme
const (
	tailleDamier  = 9
	symboleJoeur1 = "x"
	symboleJoeur2 = "0"
)

//varialbles globales qui peuvent bouger pendant la partie

var (
	tableauMorpion = [tailleDamier]string{
		"1", "2", "3",
		"4", "4", "6",
		"7", "8", "9"}

	joueur1 = true // Joueur 1 commence

)

//Afficher tableau
func affichageTableau() {

	for i := 0; i < len(tableauMorpion); i++ {

		fmt.Print(" ", tableauMorpion[i], " ")
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}

}

func SaisieSymbole() {

	scanner := bufio.NewScanner(os.Stdin)
	var saisie string
	fmt.Print("Entrez votre symbole : ")
	saisie = scanner.Scan.()
	fmt.Println(saisie)
}

//function principale

func main() {

	affichageTableau()
	SaisieSymbole()
}
