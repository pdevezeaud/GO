package main

import (
	"fmt"
	"io/ioutil"
)

//function qui lit un fichier (prend en argument un nom de fichier)
// et en retour un string et le code erreur
// ici fonction avec retour multiple
func readFile(filename string) (string, error) {
	//on stocke dans une variable
	dat, err := ioutil.ReadFile(filename)
	// si erreur est different de null alors on traite
	if err != nil {
		//ici des cotes vide car la fonction n'a pas pu lire le contrenu
		return "", err
	}

	//deuxieme test. on verifie que dat (le fichier) n'est pas à zero
	//avec la fonction errors.New on crée une nouvelle erreur.
	if len(dat) == 0 {
		//return "", errors.New("Contenu vide"). on cree une erreur avec le packag errors.New
		return "", fmt.Errorf("fichier vide (filename =%v)", filename)

	}

	// après les deux, si pas d'erreur on retourne le contenu
	return string(dat), nil

}

func main() {
	//utilisation de la fonction
	dat, err := readFile("test.txt")
	//on test si pas null
	if err != nil {
		fmt.Printf("Erreur à la lecture du fichier: %v\n", err)
		return
	}
	//sinon on affiche le contenu
	fmt.Println("Contenu du fichier: ")
	fmt.Println(dat)

}
